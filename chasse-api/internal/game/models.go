package game

import (
	"chasse-api/internal/game/engine"
	"chasse-api/internal/utils"
)

type Session struct {
	Id    string      `json:"id"`
	Mode  engine.Mode `json:"mode"`
	State any         `json:"state"`
}

type Response struct {
	Operation Operation `json:"operation"`
	Verdict   Verdict   `json:"verdict"`
	Position  string    `json:"position"`
	Session   Session   `json:"session"`
	Message   string    `json:"message"`
}

func NewResponse() *Response {
	return &Response{}
}

func (resp *Response) SetVerdict(o Verdict) *Response {
	resp.Verdict = o
	return resp
}

func (resp *Response) SetMessage(o string) *Response {
	resp.Message = o
	return resp
}

func (resp *Response) SetRequestData(o Request) *Response {
	resp.Operation = o.Operation
	return resp
}

func (resp *Response) SetSession(o Session) *Response {
	resp.Session = o
	return resp
}

func (resp *Response) SetPosition(o string) *Response {
	resp.Position = o
	return resp
}

type Request struct {
	Operation Operation   `json:"operation"`
	Move      engine.Move `json:"move"`
}

func NewRequest() *Request {
	return &Request{}
}

func (req *Request) Validate() bool {
	switch req.Operation {
	case MOVE:
		return req.Move.Target != "" && req.Move.Piece != ""
	case RESET:
		fallthrough
	case BACK:
		fallthrough
	case CLEAR:
		fallthrough
	case STATUS:
		return true
	default:
		return false
	}
}

func (req *Request) SetOperation(o Operation) *Request {
	req.Operation = o
	return req
}

func (s *Session) Encode() ([]byte, error) {
	return utils.Encode(*s)
}

func (s *Session) Decode(data []byte) error {
	return utils.Decode(data, s)
}

type Operation string

const (
	MOVE   Operation = "MOVE"
	BACK   Operation = "BACK"
	RESET  Operation = "RESET"
	CLEAR  Operation = "CLEAR"
	STATUS Operation = "STATUS"
)

type Verdict string

const (
	UPDATED Verdict = "UPDATED"
	INVALID Verdict = "INVALID"
	OK      Verdict = "OK"
)
