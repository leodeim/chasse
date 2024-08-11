package game

import (
	"chasse-api/core"
	e "chasse-api/error"
	"chasse-api/game/engine"

	"github.com/google/uuid"
)

type GameBuilder struct {
	mode    engine.Mode
	session *Session
	app     *core.App
}

func NewGameBuilder() *GameBuilder {
	gb := &GameBuilder{}
	return gb
}

func (b *GameBuilder) WithApp(app *core.App) *GameBuilder {
	b.app = app
	return b
}

func (b *GameBuilder) WithMode(mode engine.Mode) *GameBuilder {
	b.mode = mode
	return b
}

func (b *GameBuilder) WithSession(session *Session) *GameBuilder {
	b.session = session
	return b
}

func (b *GameBuilder) Build() (*Game, error) {
	if b.app == nil {
		return nil, e.BuildError(e.INTERNAL, "can't build game: core app is missing")
	}

	switch b.session {
	case nil:
		engine, err := engine.GetEngine(b.mode, nil)
		if err != nil {
			return nil, err
		}

		return &Game{
			app:       b.app,
			engine:    engine,
			sessionId: uuid.New().String(),
		}, nil
	default:
		engine, err := engine.GetEngine(b.session.Mode, b.session.State)
		if err != nil {
			return nil, err
		}

		return &Game{
			app:       b.app,
			engine:    engine,
			sessionId: b.session.Id,
		}, nil
	}
}
