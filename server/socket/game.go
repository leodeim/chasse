package socket

import (
	"fmt"
	"log"

	"github.com/leonidasdeim/zen-chess/server/models"
	"github.com/leonidasdeim/zen-chess/server/store"
)

func GameAction(data models.SessionActionMessage, client *Client, store *store.Store) {
	switch data.Action {
	case models.MOVE:
		Move(data, store)
	case models.JOIN_ROOM:
		JoinRoom(data, client, store)
	default:
		fmt.Printf("(Room %s) Bad action type: %d", data.SessionId, data.Action)
	}
}

func Move(data models.SessionActionMessage, store *store.Store) {
	room := FindRoom(data.SessionId)
	if room != nil {
		store.UpdateSession(data.SessionId, data.Position)
		room.broadcast <- &data
	}
}

func JoinRoom(data models.SessionActionMessage, client *Client, store *store.Store) {
	log.Println(data)
	if data.SessionId == "" {
		return
	}

	// verify if session is registered
	if _, err := store.GetSession(data.SessionId); err == nil {
		room := FindOrCreateRoom(data.SessionId)
		room.register <- client
	}
}
