package socket

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/leonidasdeim/zen-chess/internal/models"
	"github.com/leonidasdeim/zen-chess/internal/store"
)

type client struct{}

var clients = make(map[*websocket.Conn]client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan []byte)
var unregister = make(chan *websocket.Conn)

func Setup(app *fiber.App, store *store.Store) {
	go socketRunner()

	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		defer func() {
			unregister <- c
			c.Close()
		}()

		register <- c

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read error:", err)
				}

				return
			}

			if messageType == websocket.TextMessage {

				session := &models.SessionDataModel{}
				if err := json.Unmarshal(message, &session); err == nil {
					if _, err := store.UpdateSession(session.Uuid, session.Position); err == nil {
						broadcast <- message
					}
				}

			} else {
				log.Println("websocket message received of type", messageType)
			}
		}
	}))
}

func socketRunner() {
	for {
		select {
		case connection := <-register:
			clients[connection] = client{}
			log.Println("ws client registered")

		case message := <-broadcast:
			log.Println("ws message will be sent:", message)

			for connection := range clients {
				if err := connection.WriteMessage(websocket.TextMessage, message); err != nil {
					log.Println("ws write error:", err)

					connection.WriteMessage(websocket.CloseMessage, []byte{})
					connection.Close()
					delete(clients, connection)
				}
			}

		case connection := <-unregister:
			delete(clients, connection)
			log.Println("ws client unregistered")
		}
	}
}
