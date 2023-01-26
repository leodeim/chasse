package socket

import (
	e "chasse-api/internal/error"
	"chasse-api/internal/models"
	"log"
)

func GameAction(data models.SessionActionMessage, h *SocketHandler) error {
	switch data.Action {
	case models.MOVE:
		return Move(data, h)
	case models.JOIN_ROOM:
		return JoinRoom(data, h)
	default:
		log.Printf("(Room %s) Bad action type: %d", data.SessionId, data.Action)
		return e.BadRequest{Message: "bad action type"}
	}
}

func Move(data models.SessionActionMessage, h *SocketHandler) error {
	room := FindRoom(data.SessionId)
	if room != nil {
		_, err := h.store.UpdateSession(data.SessionId, data.Position)
		if err != nil {
			return err
		}
		room.broadcast <- &BroadcastData{
			message: &data,
			client:  h.client,
		}
	} else {
		return e.BadRequest{Message: "room has not been found"}
	}

	return nil
}

func JoinRoom(data models.SessionActionMessage, h *SocketHandler) error {
	if data.SessionId == "" {
		return e.BadRequest{Message: "sessionId is empty"}
	}

	// verify if session is registered
	if session, err := h.store.GetSession(data.SessionId); err == nil {
		room := FindOrCreateRoom(data.SessionId)
		room.register <- &BroadcastData{
			message: session,
			client:  h.client,
		}
	} else {
		return err
	}

	return nil
}
