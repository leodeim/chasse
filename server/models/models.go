package models

import (
	"encoding/json"
	"log"
)

type WebsocketAction int

const (
	MOVE      WebsocketAction = 0
	GO_BACK   WebsocketAction = 1
	RESET     WebsocketAction = 2
	JOIN_ROOM WebsocketAction = 3
	ERROR     WebsocketAction = 4
	OK        WebsocketAction = 5
)

type SessionActionMessage struct {
	Action    WebsocketAction `json:"action"`
	SessionId string          `json:"sessionId"`
	Position  string          `json:"position"`
}

func ErrorMessage() SessionActionMessage {
	return SessionActionMessage{
		Action: ERROR,
	}
}

func OkMessage() SessionActionMessage {
	return SessionActionMessage{
		Action: OK,
	}
}

func (message *SessionActionMessage) Encode() []byte {
	json, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
	}

	return json
}

var BLANK_BOARD = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1"
