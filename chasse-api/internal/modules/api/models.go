package api

import (
	"chasse-api/internal/game/engine"
)

type NewSessionRequest struct {
	Mode engine.Mode `json:"mode" validate:"required"`
}
