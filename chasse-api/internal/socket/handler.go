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

type Client struct {
	conn      *websocket.Conn
	sessionId string
	notifier  *gobrake.Notifier
}

type SocketHandler struct {
	client *Client
	store  *store.Store
}

func InitClient(app *fiber.App, s *store.Store, n *gobrake.Notifier) {
	app.Get("/api/ws", websocket.New(func(c *websocket.Conn) {
		h := &SocketHandler{
			client: &Client{
				conn:     c,
				notifier: n,
			},
			store: s,
		}

		log.Printf("(Client %s) Logged in\n", h.client.conn.RemoteAddr())
		h.respondOk(models.CONNECT)

		defer func() {
			log.Printf("(Client %s) Logged out\n", h.client.conn.RemoteAddr())
			RemoveClientFromRoom(h.client)
			h.client.conn.Close()
		}()

		for {
			messageType, rawMessage, err := h.client.conn.ReadMessage()
			if err != nil {
				log.Printf("(Client %s) WebSocket client read error: %v\n", h.client.conn.RemoteAddr(), err)
				return
			}

			if messageType == websocket.TextMessage {
				message := &models.SessionActionMessage{}
				if err := json.Unmarshal(rawMessage, &message); err != nil {
					h.respondError(models.BLANK_ACTION, err)
					continue
				}
				if err := GameAction(*message, h); err != nil {
					h.respondError(message.Action, err)
					continue
				}
			} else {
				log.Printf("(Client %s) WebSocket message received of type: %d\n", h.client.conn.RemoteAddr(), messageType)
			}
		}
	}))
}

func (h *SocketHandler) respondOk(action models.WebsocketAction) {
	msg := models.OkMessage(action)
	h.client.conn.WriteMessage(websocket.TextMessage, msg.Encode())
}

func (h *SocketHandler) respondError(action models.WebsocketAction, err error) {
	h.client.notifier.Notify(err, nil) // send error to airbrake
	errorMsg := models.ErrorMessage(action)
	h.client.conn.WriteMessage(websocket.TextMessage, errorMsg.Encode())
	log.Printf("(Client %s) Error: %s \n", h.client.conn.LocalAddr(), err.Error())
}
