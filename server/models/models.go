package models

import (
	"encoding/json"
	"log"
)

type WebsocketAction int

const (
	BLANK_ACTION WebsocketAction = 0
	MOVE         WebsocketAction = 1
	GO_BACK      WebsocketAction = 2
	RESET        WebsocketAction = 3
	JOIN_ROOM    WebsocketAction = 4
)

type WebsocketResponse int

const (
	BLANK_RESPONSE WebsocketResponse = 0
	OK             WebsocketResponse = 1
	ERROR          WebsocketResponse = 2
)

type SessionActionMessage struct {
	Action    WebsocketAction   `json:"action"`
	Response  WebsocketResponse `json:"response"`
	SessionId string            `json:"sessionId"`
	Position  string            `json:"position"`
}

func ErrorMessage(action WebsocketAction) SessionActionMessage {
	return SessionActionMessage{
		Action:   action,
		Response: ERROR,
	}
}

func OkMessage(action WebsocketAction) SessionActionMessage {
	return SessionActionMessage{
		Action:   action,
		Response: OK,
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
