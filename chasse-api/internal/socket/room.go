package socket

import (
	"fmt"
	"log"

	"chasse-api/internal/models"

	"github.com/gofiber/websocket/v2"
)

type Room struct {
	SessionId  string `json:"id"`
	clients    map[*Client]bool
	register   chan *BroadcastData
	unregister chan *Client
	broadcast  chan *BroadcastData
}

type BroadcastData struct {
	message *models.SessionActionMessage
	client  *Client
}

var activeRooms = make(map[string]*Room)

func FindOrCreateRoom(id string) *Room {
	room := FindRoom(id)

	if room == nil {
		room = &Room{
			SessionId:  id,
			clients:    make(map[*Client]bool),
			register:   make(chan *BroadcastData),
			unregister: make(chan *Client),
			broadcast:  make(chan *BroadcastData),
		}

		go room.runner()
	}

	return room
}

func RemoveClientFromRoom(client *Client) {
	if room := FindRoom(client.sessionId); room != nil {
		room.unregister <- client
	}
}

func FindRoom(id string) *Room {
	if room, found := activeRooms[id]; found {
		return room
	}
	return nil
}

func (room *Room) runner() {
	fmt.Printf("(Room %s) Runner is starting \n", room.SessionId)
	activeRooms[room.SessionId] = room

	defer func() {
		fmt.Printf("(Room %s) Runner is stopping \n", room.SessionId)
		delete(activeRooms, room.SessionId)
	}()

	for {
		select {
		case data := <-room.register:
			message := models.SessionActionMessage{
				Action:    models.MOVE,
				SessionId: data.message.SessionId,
				Position:  data.message.Position,
			}
			if err := data.client.conn.WriteMessage(websocket.TextMessage, message.Encode()); err == nil {
				data.client.sessionId = room.SessionId
				room.clients[data.client] = true
				log.Printf("(Room %s) Client registered, clients in the room: %d \n", room.SessionId, len(room.clients))
			}

		case data := <-room.broadcast:
			for client := range room.clients {
				if client == data.client {
					continue
				}
				if err := client.conn.WriteMessage(websocket.TextMessage, data.message.Encode()); err != nil {
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
