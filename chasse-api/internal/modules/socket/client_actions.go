package socket

import e "chasse-api/internal/error"

func handleClientMessage(c *Client, msg *Message) error {
	log.Debugf("(client %s) received message: %v\n", c.connection.RemoteAddr(), msg.String())

	switch msg.Type {
	case GAME:
		return sendAction(c, msg)
	case JOIN_ROOM:
		return joinRoom(c, msg)
	case LEAVE_ROOM:
		return leaveRoom(c, msg)
	case PING:
		c.respond(response{msg, OK})
		return nil
	default:
		return e.BuildErrorf(e.BAD_REQUEST, "bad message type: %s", msg.String())
	}
}

func sendAction(c *Client, msg *Message) error {
	if room := findRoom(c.sessionId); room != nil {
		room.actions <- &BroadcastData{
			message: msg,
			client:  c,
		}
	} else {
		return e.BuildError(e.BAD_REQUEST, "room not found")
	}

	return nil
}

func joinRoom(c *Client, msg *Message) error {
	if msg.Request.SessionId == "" {
		return e.BuildError(e.BAD_REQUEST, "session id is empty")
	}

	room, err := findOrCreateRoom(c.app, msg.Request.SessionId)
	if err != nil {
		return err
	}

	room.register <- &BroadcastData{
		message: msg,
		client:  c,
	}

	return nil
}

func leaveRoom(c *Client, msg *Message) error {
	if room := findRoom(c.sessionId); room != nil {
		room.unregister <- &BroadcastData{
			message: msg,
			client:  c,
		}
		return nil
	}

	return e.BuildError(e.BAD_REQUEST, "leave room: room not found")
}
