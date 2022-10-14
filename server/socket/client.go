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

		respondOk := func(action models.WebsocketAction) {
			msg := models.OkMessage(action)
			client.conn.WriteMessage(websocket.TextMessage, msg.Encode())
		}
		respondError := func(action models.WebsocketAction, err error) {
			errorMsg := models.ErrorMessage(action)
			client.conn.WriteMessage(websocket.TextMessage, errorMsg.Encode())
			log.Printf("(Client %s) Error: %s \n", client.conn.LocalAddr(), err.Error())
		}

		for {
			messageType, rawMessage, err := c.ReadMessage()
			if err != nil {
				log.Printf("(Client %s) WebSocket client read error: %v \n", client.conn.LocalAddr(), err)
				return
			}

			if messageType == websocket.TextMessage {
				message := &models.SessionActionMessage{}

				if err := json.Unmarshal(rawMessage, &message); err != nil {
					respondError(models.BLANK_ACTION, err)
					continue
				}

				if err := GameAction(*message, client, store); err != nil {
					respondError(message.Action, err)
					continue
				} else {
					respondOk(message.Action)
				}

			} else {
				log.Printf("(Client %s) WebSocket message received of type: %d \n", client.conn.LocalAddr(), messageType)
			}
		}
	}))
}
