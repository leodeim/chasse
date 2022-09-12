package socket

import (
	"fmt"
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/leonidasdeim/zen-chess/server/models"
)

type Room struct {
	SessionId  string `json:"id"`
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan *models.SessionDataModel
}

var rooms = make(map[*Room]bool)

func CreateOrGetRoom(id string) *Room {
	room := findExistingRoom(id)

	if room == nil {
		room = &Room{
			SessionId:  id,
			clients:    make(map[*Client]bool),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			broadcast:  make(chan *models.SessionDataModel),
		}

		rooms[room] = true
		go room.roomRunner()
	}

	return room
}

func findExistingRoom(id string) *Room {
	for room := range rooms {
		if room.SessionId == id {
			return room
		}
	}

	return nil
}

func (room *Room) roomRunner() {
	fmt.Printf("(Room %s) Runner starting \n", room.SessionId)

	defer func() {
		fmt.Printf("(Room %s) Runner stopped \n", room.SessionId)
		delete(rooms, room)
	}()

	for {
		select {
		case client := <-room.register:
			room.clients[client] = true
			log.Printf("(Room %s) Client registered \n", room.SessionId)

		case message := <-room.broadcast:
			log.Printf("(Room %s) Message will be sent: %+v\n", room.SessionId, message)

			for client := range room.clients {
				if err := client.conn.WriteMessage(websocket.TextMessage, message.Encode()); err != nil {
					log.Printf("(Room %s) WebSocket write error: %v", room.SessionId, err)

					client.conn.WriteMessage(websocket.CloseMessage, []byte{})
					client.conn.Close()
					delete(room.clients, client)
				}
			}

		case client := <-room.unregister:
			delete(room.clients, client)
			log.Printf("(Room %s) Client unregistered \n", room.SessionId)
		}
	}
}
