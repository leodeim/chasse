package socket

import (
	"fmt"
	"log"

	e "chasse-api/internal/error"
	"chasse-api/internal/models"
	"chasse-api/internal/store"
)

func GameAction(data models.SessionActionMessage, client *Client, store *store.Store) error {
	switch data.Action {
	case models.MOVE:
		return Move(data, client, store)
	case models.JOIN_ROOM:
		return JoinRoom(data, client, store)
	default:
		fmt.Printf("(Room %s) Bad action type: %d", data.SessionId, data.Action)
		return e.BadRequest{Message: "bad action type"}
	}
}

func Move(data models.SessionActionMessage, client *Client, store *store.Store) error {
	room := FindRoom(data.SessionId)
	if room != nil {
		_, err := store.UpdateSession(data.SessionId, data.Position)
		if err != nil {
			return err
		}
		room.broadcast <- &BroadcastData{
			message: &data,
			client:  client,
		}
	} else {
		return e.BadRequest{Message: "room has not been found"}
	}

	return nil
}

func JoinRoom(data models.SessionActionMessage, client *Client, store *store.Store) error {
	log.Println(data)
	if data.SessionId == "" {
		return e.BadRequest{Message: "sessionId is empty"}
	}

	// verify if session is registered
	if session, err := store.GetSession(data.SessionId); err == nil {
		room := FindOrCreateRoom(data.SessionId)
		room.register <- &BroadcastData{
			message: session,
			client:  client,
		}
	} else {
		return err
	}

	return nil
}
