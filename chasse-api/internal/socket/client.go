package socket

import (
	"encoding/json"
	"log"

	"chasse-api/internal/models"
	"chasse-api/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	conn      *websocket.Conn
	store     *store.Store
	sessionId string
}

func InitClient(app *fiber.App, store *store.Store) {
	app.Get("/api/ws", websocket.New(func(conn *websocket.Conn) {
		c := &Client{
			conn:  conn,
			store: store,
		}

		log.Printf("(Client %s) Logged in\n", c.conn.RemoteAddr())
		c.respondOk(models.CONNECT)

		defer func() {
			log.Printf("(Client %s) Logged out\n", c.conn.RemoteAddr())
			RemoveClientFromRoom(c)
			c.conn.Close()
		}()

		for {
			messageType, rawMessage, err := c.conn.ReadMessage()
			if err != nil {
				log.Printf("(Client %s) WebSocket client read error: %v\n", c.conn.RemoteAddr(), err)
				return
			}

			if messageType == websocket.TextMessage {
				message := &models.SessionActionMessage{}
				if err := json.Unmarshal(rawMessage, &message); err != nil {
					c.respondError(models.BLANK_ACTION, err)
					continue
				}
				if err := GameAction(*message, c, store); err != nil {
					c.respondError(message.Action, err)
					continue
				}
			} else {
				log.Printf("(Client %s) WebSocket message received of type: %d\n", c.conn.RemoteAddr(), messageType)
			}
		}
	}))
}

func (c *Client) respondOk(action models.WebsocketAction) {
	msg := models.OkMessage(action)
	c.conn.WriteMessage(websocket.TextMessage, msg.Encode())
}

func (c *Client) respondError(action models.WebsocketAction, err error) {
	errorMsg := models.ErrorMessage(action)
	c.conn.WriteMessage(websocket.TextMessage, errorMsg.Encode())
	log.Printf("(Client %s) Error: %s \n", c.conn.LocalAddr(), err.Error())
}
