package engine

import (
	e "chasse-api/error"
	"encoding/json"
)

type free struct {
	positions []string
}

const maxPositions = 50

var _ Engine = (*free)(nil)

type FreeState struct {
	Positions []string `json:"position"`
}

func InitFreeMode(state ...*FreeState) *free {
	f := &free{}
	if len(state) > 0 && state[0] != nil {
		f.addPositions(state[0].Positions...)
	} else {
		f.addPositions(BLANK_BOARD)
	}

	return f
}

func (f *free) addPositions(pos ...string) {
	if len(f.positions) >= maxPositions {
		f.positions = f.positions[1:]
	}
	f.positions = append(f.positions, pos...)
}

func (f *free) Move(move Move) (bool, error) {
	posMap := make(map[string]string)
	if err := json.Unmarshal([]byte(f.GetPosition()), &posMap); err != nil {
		return false, e.BuildErrorf(e.INTERNAL, "failed while unmarshal position object: %s", err.Error())
	}

	if move.Source == move.Target && posMap[move.Target] == move.Piece {
		return false, nil
	}

	if move.Source == "spare" && move.Target == "offBoard" {
		return false, nil
	}

	if v, ok := posMap[move.Source]; !ok || v != move.Piece {
		return false, nil
	}

	if move.Target != "offBoard" {
		posMap[move.Target] = move.Piece
	}
	delete(posMap, move.Source)

	bytes, err := json.Marshal(posMap)
	if err != nil {
		return false, e.BuildErrorf(e.INTERNAL, "failed while marshal position object: %s", err.Error())
	}

	f.addPositions(string(bytes))

	return true, nil
}

func (f *free) Back() (bool, error) {
	if len(f.positions) <= 1 {
		return false, e.BuildErrorf(e.BAD_REQUEST, "nowhere to go back")
	}

	f.positions = f.positions[:len(f.positions)-1]
	return true, nil
}

func (f *free) Reset() (bool, error) {
	f.addPositions(BLANK_BOARD)
	return true, nil
}

func (f *free) Clear() (bool, error) {
	f.addPositions(CLEAR_BOARD)
	return true, nil
}

func (f *free) GetState() any {
	return FreeState{
		Positions: f.positions,
	}
}

func (f *free) GetPosition() string {
	if l := len(f.positions); l > 0 {
		return f.positions[l-1]
	}
	return ""
}

func (*free) GetMode() Mode {
	return FREE_MODE
}
