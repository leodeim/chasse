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
	conn *websocket.Conn
}

func serveClient(app *fiber.App, store *store.Store) {
	app.Get("/ws/:session", websocket.New(func(c *websocket.Conn) {
		sessionId := c.Params("session")
		room := CreateOrGetRoom(sessionId)
		client := &Client{conn: c}
		room.register <- client

		defer func() {
			room.unregister <- client
			c.Close()
		}()

		for {
			messageType, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("Read ERROR:", err)
				}

				return
			}

			if messageType == websocket.TextMessage {

				session := &models.SessionDataModel{}
				if err := json.Unmarshal(message, &session); err == nil {
					if _, err := store.UpdateSession(session.SessionId, session.Position); err == nil {
						room.broadcast <- session
					}
				}

			} else {
				log.Println("WebSocket message received of type", messageType)
			}
		}
	}))
}
