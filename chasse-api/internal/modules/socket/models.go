package socket

import (
	"chasse-api/internal/game"
	"chasse-api/internal/utils"
)

type Request struct {
	SessionId string       `json:"sessionId"`
	Game      game.Request `json:"game"`
}

type Response struct {
	Status ResponseCode  `json:"status"`
	Game   game.Response `json:"game"`
}

type Message struct {
	Type     MessageType `json:"type"`
	Request  Request     `json:"request"`
	Response Response    `json:"response"`
}

func (m *Message) addResponseStatus(status ResponseCode) *Message {
	m.Response.Status = status
	return m
}

func (m *Message) addGameResponse(game game.Response) *Message {
	m.Response.Game = game
	return m
}

type MessageType string

const (
	BLANK_TYPE MessageType = "BLANK"
	CONNECT    MessageType = "CONNECT"
	JOIN_ROOM  MessageType = "JOIN_ROOM"
	LEAVE_ROOM MessageType = "LEAVE_ROOM"
	GAME       MessageType = "GAME"
	PING       MessageType = "PING"
)

type ResponseCode string

const (
	BLANK_RESPONSE ResponseCode = "BLANK"
	OK             ResponseCode = "OK"
	ERROR          ResponseCode = "ERROR"
)

func (message *Message) Encode() ([]byte, error) {
	return utils.Encode(*message)
}

func (message *Message) Decode(data []byte) error {
	return utils.Decode(data, message)
}

func (message *Message) String() string {
	data, err := message.Encode()
	if err != nil {
		return "can't marshall message to a string"
	}
	return string(data)
}

type BroadcastData struct {
	message *Message
	client  *Client
}

func validateBroadcastData(data *BroadcastData) bool {
	if data == nil || data.message == nil || data.client == nil {
		log.Error("failed to validate BroadcastData object")
		return false
	}
	return true
}
