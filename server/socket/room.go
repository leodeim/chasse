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

var activeRooms = make(map[*Room]bool)

func FindOrCreateRoom(id string) *Room {
	// TODO: verify room in redis
	room := findActiveRoom(id)

	if room == nil {
		room = &Room{
			SessionId:  id,
			clients:    make(map[*Client]bool),
			register:   make(chan *Client),
			unregister: make(chan *Client),
			broadcast:  make(chan *models.SessionDataModel),
		}

		go room.run()
	}

	return room
}

func findActiveRoom(id string) *Room {
	for room := range activeRooms {
		if room.SessionId == id {
			return room
		}
	}

	return nil
}

func (room *Room) run() {
	fmt.Printf("(Room %s) Runner is starting \n", room.SessionId)
	activeRooms[room] = true

	defer func() {
		fmt.Printf("(Room %s) Runner is stopping \n", room.SessionId)
		delete(activeRooms, room)
	}()

	for {
		select {
		case client := <-room.register:
			room.clients[client] = true
			log.Printf("(Room %s) Client registered, clients in the room: %d \n", room.SessionId, len(room.clients))

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
			log.Printf("(Room %s) Client unregistered, clients in the room: %d \n", room.SessionId, len(room.clients))

			if len(room.clients) < 1 {
				log.Printf("(Room %s) Is empty \n", room.SessionId)
				return
			}
		}
	}
}
