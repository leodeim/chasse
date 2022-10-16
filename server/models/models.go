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

var BLANK_BOARD = `{"a8":"bR","b8":"bN","c8":"bB","d8":"bQ","e8":"bK","f8":"bB","g8":"bN","h8":"bR","a7":"bP","b7":"bP","c7":"bP","d7":"bP","e7":"bP","f7":"bP","g7":"bP","h7":"bP","a2":"wP","b2":"wP","c2":"wP","e2":"wP","f2":"wP","h2":"wP","a1":"wR","b1":"wN","c1":"wB","d1":"wQ","e1":"wK","f1":"wB","g1":"wN","h1":"wR","g2":"wP","d2":"wP"}`
