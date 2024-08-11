package game

import (
	"chasse-api/core"
	"chasse-api/game/engine"
)

// var log = logger.New("GAME")

const notImplResp = "not implemented"

type Game struct {
	app       *core.App
	engine    engine.Engine
	sessionId string
}

func New(app *core.App, mode engine.Mode) (*Game, error) {
	game, err := NewGameBuilder().
		WithApp(app).
		WithMode(mode).
		Build()
	if err != nil {
		return nil, err
	}

	if err = game.store(); err != nil {
		return nil, err
	}

	return game, nil
}

func Load(app *core.App, sessionId string) (*Game, error) {
	s := &Session{}
	data, err := app.Store.Read(sessionId)
	if err != nil {
		return nil, err
	}

	if err = s.Decode(data); err != nil {
		return nil, err
	}

	return NewGameBuilder().
		WithApp(app).
		WithSession(s).
		Build()
}

func (g *Game) store() error {
	session := &Session{
		Id:    g.sessionId,
		Mode:  g.engine.GetMode(),
		State: g.engine.GetState(),
	}

	data, err := session.Encode()
	if err != nil {
		return err
	}

	if err = g.app.Store.Write(g.sessionId, data); err != nil {
		return err
	}

	return nil
}

func (g *Game) Do(req *Request) *Response {
	if !req.Validate() {
		return NewResponse().
			SetRequestData(*req).
			SetVerdict(INVALID).
			SetMessage("failed to validate game request")
	}

	resp := func() *Response {
		switch req.Operation {
		case MOVE:
			return g.move(req)
		case BACK:
			return g.back(req)
		case RESET:
			return g.reset(req)
		case CLEAR:
			return g.clear(req)
		case STATUS:
			return g.getStatus(req)
		default:
			return NewResponse().
				SetRequestData(*req).
				SetVerdict(INVALID).
				SetMessage("bad operation")
		}
	}()

	if resp != nil && resp.Verdict == UPDATED {
		g.store()
	}
	return resp
}

func (g *Game) move(req *Request) *Response {
	ok, err := g.engine.Move(req.Move)

	resp := NewResponse().
		SetRequestData(*req).
		SetPosition(g.engine.GetPosition())

	if err != nil {
		return resp.SetVerdict(INVALID).SetMessage(err.Error())
	} else if !ok {
		return resp.SetVerdict(INVALID)
	}

	return resp.SetVerdict(UPDATED)
}

func (g *Game) back(req *Request) *Response {
	return NewResponse().SetRequestData(*req).SetVerdict(INVALID).SetMessage(notImplResp)
}

func (g *Game) reset(req *Request) *Response {
	return NewResponse().SetRequestData(*req).SetVerdict(INVALID).SetMessage(notImplResp)
}

func (g *Game) clear(req *Request) *Response {
	return NewResponse().SetRequestData(*req).SetVerdict(INVALID).SetMessage(notImplResp)
}

func (g *Game) getStatus(req *Request) *Response {
	return NewResponse().
		SetRequestData(*req).
		SetPosition(g.engine.GetPosition()).
		SetSession(Session{g.sessionId, g.engine.GetMode(), g.engine.GetState()}).
		SetVerdict(OK)
}
