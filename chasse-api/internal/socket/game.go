package socket

import (
	"fmt"
	"log"

	e "chasse-api/internal/error"
	"chasse-api/internal/models"
)

func GameAction(data models.SessionActionMessage, c *ClientHandler) error {
	switch data.Action {
	case models.MOVE:
		return Move(data, c)
	case models.JOIN_ROOM:
		return JoinRoom(data, c)
	default:
		fmt.Printf("(Room %s) Bad action type: %d", data.SessionId, data.Action)
		return e.BadRequest{Message: "bad action type"}
	}
}

func Move(data models.SessionActionMessage, c *ClientHandler) error {
	room := FindRoom(data.SessionId)
	if room != nil {
		_, err := c.store.UpdateSession(data.SessionId, data.Position)
		if err != nil {
			return err
		}
		room.broadcast <- &BroadcastData{
			message: &data,
			client:  c,
		}
	} else {
		return e.BadRequest{Message: "room has not been found"}
	}

	return nil
}

func JoinRoom(data models.SessionActionMessage, c *ClientHandler) error {
	log.Println(data)
	if data.SessionId == "" {
		return e.BadRequest{Message: "sessionId is empty"}
	}

	// verify if session is registered
	if session, err := c.store.GetSession(data.SessionId); err == nil {
		room := FindOrCreateRoom(data.SessionId)
		room.register <- &BroadcastData{
			message: session,
			client:  c,
		}
	} else {
		return err
	}

	return nil
}
