package api

import (
	"chasse-api/game/engine"
)

type NewSessionRequest struct {
	Mode engine.Mode `json:"mode" validate:"required"`
}
