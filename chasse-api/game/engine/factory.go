package engine

import (
	e "chasse-api/error"
	"chasse-api/utils"
)

type Mode string

const (
	FREE_MODE   Mode = "FREE"
	STRICT_MODE Mode = "STRICT"
)

type Move struct {
	Source string `json:"source"`
	Target string `json:"target"`
	Piece  string `json:"piece"`
}

type Engine interface {
	Move(Move) (bool, error)
	GetState() any
	GetPosition() string
	GetMode() Mode
}

func GetEngine(mode Mode, state any) (Engine, error) {
	switch mode {
	case FREE_MODE:
		fs, err := resolveFreeState(state)
		if err != nil {
			return nil, err
		}
		return InitFreeMode(fs), nil

	case STRICT_MODE:
		ss, err := resolveStrictState(state)
		if err != nil {
			return nil, err
		}
		return InitStrictMode(ss), nil
	default:

		return nil, e.BuildError(e.BAD_REQUEST, "bad engine mode")
	}
}

func resolveFreeState(s any) (*FreeState, error) {
	if s == nil {
		return nil, nil
	}

	fs, err := utils.MapToStruct[FreeState](s)
	if fs == nil || err != nil || len(fs.Positions) == 0 {
		return nil, e.BuildError(e.INTERNAL, "bad free mode engine state")
	}

	return fs, nil
}

func resolveStrictState(s any) (*StrictState, error) {
	if s == nil {
		return nil, nil
	}

	ss, err := utils.MapToStruct[StrictState](s)
	if ss == nil || err != nil {
		return nil, e.BuildError(e.INTERNAL, "bad strict mode engine state")
	}

	return ss, nil
}
