package socket

import (
	"chasse-api/core"
	e "chasse-api/error"
	"chasse-api/game"
)

type Room struct {
	sessionId  string
	clients    map[*Client]bool
	game       *game.Game
	register   chan *BroadcastData
	unregister chan *BroadcastData
	actions    chan *BroadcastData
}

var activeRooms = make(map[string]*Room)

func findOrCreateRoom(app *core.App, id string) (*Room, error) {
	if room := findRoom(id); room != nil {
		return room, nil
	}

	game, err := game.Load(app, id)
	if err != nil {
		return nil, e.BuildErrorf(e.BAD_REQUEST, "can't load game instance: %s, reason: %s", id, err.Error())
	}

	room := &Room{
		sessionId:  id,
		clients:    make(map[*Client]bool),
		game:       game,
		register:   make(chan *BroadcastData),
		unregister: make(chan *BroadcastData),
		actions:    make(chan *BroadcastData),
	}

	go room.runner()

	return room, nil
}

func findRoom(id string) *Room {
	if room, ok := activeRooms[id]; ok {
		return room
	}
	return nil
}

func (room *Room) runner() {
	// prevent running same room in different goroutine
	if r := findRoom(room.sessionId); r != nil {
		return
	}

	activeRooms[room.sessionId] = room
	log.Infof("(room %s) runner is starting \n", room.sessionId)

	defer func() {
		log.Infof("(room %s) runner is stopping \n", room.sessionId)
		delete(activeRooms, room.sessionId)
	}()

	for {
		select {
		case msg := <-room.actions:
			room.handleAction(msg)
		case msg := <-room.register:
			room.handleRegister(msg)
		case msg := <-room.unregister:
			if shouldStop := room.handleUnregister(msg); shouldStop {
				return
			}
		}
	}
}

func (room *Room) handleAction(data *BroadcastData) {
	if !validateBroadcastData(data) {
		return
	}

	gameResponse := room.game.Do(&data.message.Request.Game)

	for client := range room.clients {
		if err := client.respond(response{data.message.addGameResponse(*gameResponse), OK}); err != nil {
			log.Errorf("(room %s) handleAction: %v", room.sessionId, err)
			client.app.Monitor.Notify(err)
			client.close()

			room.unregister <- data
		}
	}
}

func (room *Room) handleRegister(data *BroadcastData) {
	if !validateBroadcastData(data) {
		return
	}

	gameResponse := room.game.Do(game.NewRequest().SetOperation(game.STATUS))

	client := data.client
	if err := client.respond(response{data.message.addGameResponse(*gameResponse), OK}); err != nil {
		log.Errorf("(room %s) handleRegister: %v", room.sessionId, err)
		client.app.Monitor.Notify(err)
		client.close()
	} else {
		client.setSessionId(room.sessionId)
		room.clients[client] = true
		log.Infof("(room %s) clients in the room: %d", room.sessionId, len(room.clients))
	}
}

func (room *Room) handleUnregister(data *BroadcastData) bool {
	if !validateBroadcastData(data) {
		return false
	}

	client := data.client
	delete(room.clients, client)
	client.respond(response{data.message, OK})

	log.Infof("(room %s) client unregistered, clients in the room: %d \n", room.sessionId, len(room.clients))

	if len(room.clients) < 1 {
		log.Infof("(room %s) is empty \n", room.sessionId)
		return true
	}
	return false
}
