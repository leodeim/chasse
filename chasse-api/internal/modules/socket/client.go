package socket

import (
	"chasse-api/internal/core"
	e "chasse-api/internal/error"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

type Client struct {
	app        *core.App
	connection *websocket.Conn
	sessionId  string
}

func handleClient(app *core.App) func(*fiber.Ctx) error {
	return websocket.New(func(c *websocket.Conn) {
		client := &Client{
			app:        app,
			connection: c,
		}

		log.Infof("(%s) logged in\n", client.connection.RemoteAddr())
		client.respond(response{&Message{Type: CONNECT}, OK})

		defer func() {
			log.Infof("(%s) logged out\n", client.connection.RemoteAddr())
			leaveRoom(client, &Message{Type: LEAVE_ROOM})
			client.connection.Close()
		}()

		client.runner()
	})
}

func (client *Client) runner() {
	log.Debugf("(%s) client runner started", client.connection.RemoteAddr())
	defer log.Debugf("(%s) client runner stopped", client.connection.RemoteAddr())

	for {
		messageType, data, err := client.connection.ReadMessage()
		if err != nil {
			log.Errorf("(%s) websocket client read error: %v\n", client.connection.RemoteAddr(), err)
			return
		}

		if messageType != websocket.TextMessage {
			log.Debugf("(%s) wrong websocket message type: %d\n", client.connection.RemoteAddr(), messageType)
			continue
		}

		message := &Message{}
		if err = message.Decode(data); err != nil {
			log.Debugf("(%s) can't decode received message: %v\n", client.connection.RemoteAddr(), data)
			client.error(err, response{message, ERROR})
			continue
		}

		if err = handleClientMessage(client, message); err != nil {
			client.error(err, response{message, ERROR})
			continue
		}
	}
}

func (client *Client) setSessionId(id string) {
	client.sessionId = id
}

type response struct {
	message *Message
	status  ResponseCode
}

func (client *Client) respond(resp response) error {
	if resp.message == nil {
		log.Error("respond error: message == nil")
		return nil
	}

	if resp.status == "" {
		log.Error("respond error: response status is empty")
		return nil
	}

	data, err := resp.message.addResponseStatus(resp.status).Encode()
	if err != nil {
		log.Errorf("respond error: can't encode message: %v", err)
		return nil
	}

	log.Debugf("(%s) response message: %s \n", client.connection.RemoteAddr(), string(data))

	if err := client.connection.WriteMessage(websocket.TextMessage, data); err != nil {
		return e.BuildErrorf(e.INTERNAL, "respond error: closing client, can't write to websocket: %v", err)
	}

	return nil
}

func (client *Client) error(err error, resp response) {
	client.app.Monitor.Notify(err)
	log.Errorf("(%s) error: %s \n", client.connection.RemoteAddr(), err.Error())

	client.respond(resp)

}

func (client *Client) close() {
	client.connection.WriteMessage(websocket.CloseMessage, []byte{})
	client.connection.Close()
}
