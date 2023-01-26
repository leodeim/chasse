package socket

import (
	"encoding/json"
	"log"

	"chasse-api/internal/models"
	"chasse-api/internal/store"

	"github.com/airbrake/gobrake/v5"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type ClientHandler struct {
	conn      *websocket.Conn
	sessionId string
	notifier  *gobrake.Notifier
	store     *store.Store
}

func InitClient(app *fiber.App, s *store.Store, n *gobrake.Notifier) {
	app.Get("/api/ws", websocket.New(func(conn *websocket.Conn) {
		c := &ClientHandler{
			conn:     conn,
			notifier: n,
			store:    s,
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
				if err := GameAction(*message, c); err != nil {
					c.respondError(message.Action, err)
					continue
				}
			} else {
				log.Printf("(Client %s) WebSocket message received of type: %d\n", c.conn.RemoteAddr(), messageType)
			}
		}
	}))
}

func (c *ClientHandler) respondOk(action models.WebsocketAction) {
	msg := models.OkMessage(action)
	c.conn.WriteMessage(websocket.TextMessage, msg.Encode())
}

func (c *ClientHandler) respondError(action models.WebsocketAction, err error) {
	c.notifier.Notify(err, nil) // send error to airbrake
	errorMsg := models.ErrorMessage(action)
	c.conn.WriteMessage(websocket.TextMessage, errorMsg.Encode())
	log.Printf("(Client %s) Error: %s \n", c.conn.LocalAddr(), err.Error())
}
