package socket

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/leonidasdeim/zen-chess/server/models"
	"github.com/leonidasdeim/zen-chess/server/store"
)

type Client struct {
	conn      *websocket.Conn
	sessionId string
}

func serveClient(app *fiber.App, store *store.Store) {
	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
		client := &Client{conn: c}
		log.Printf("(Client %s) Logged in\n", client.conn.LocalAddr())

		defer func() {
			RemoveClientFromRoom(client)
			c.Close()
		}()

		for {
			messageType, rawMessage, err := c.ReadMessage()
			if err != nil {
				log.Printf("(Client %s) WebSocket client read error: %v \n", client.conn.LocalAddr(), err)
				return
			}

			if messageType == websocket.TextMessage {
				message := &models.SessionActionMessage{}
				if err := json.Unmarshal(rawMessage, &message); err == nil {
					GameAction(*message, client, store)
				} else {
					log.Printf("(Client %s) Cant decode WebSocket message \n", client.conn.LocalAddr())
				}
			} else {
				log.Printf("(Client %s) WebSocket message received of type: %d \n", client.conn.LocalAddr(), messageType)
			}
		}
	}))
}
