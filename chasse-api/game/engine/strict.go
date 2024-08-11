package engine

// TODO: implement regular mode
type strict struct {
}

var _ Engine = (*strict)(nil)

type StrictState struct {
}

func InitStrictMode(state ...*StrictState) *strict {
	// pgn, err := chess.PGN(strings.NewReader(game.session.PGN))
	// if err != nil {
	// 	return nil, e.BuildError(e.INTERNAL, "failed while read pgn", err.Error())
	// }
	return &strict{}
}

func (e *strict) Move(move Move) (bool, error) {
	return false, nil
}

func (e *strict) Back() (bool, error) {
	return false, nil
}

func (e *strict) Reset() (bool, error) {
	return false, nil
}

func (e *strict) Clear() (bool, error) {
	return false, nil
}

func (e *strict) GetState() any {
	return "TODO"
}

func (e *strict) GetPosition() string {
	return "TODO"
}

func (e *strict) GetMode() Mode {
	return STRICT_MODE
}
