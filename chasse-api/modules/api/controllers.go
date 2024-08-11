package api

import (
	"net/http"

	e "chasse-api/error"
	"chasse-api/game"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (api *Api) HealthCheck(c *fiber.Ctx) error {
	payload := struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Config  string `json:"config"`
	}{
		api.app.Config.Get().AppName,
		api.app.Config.Get().Version,
		api.app.Config.GetTimestamp(),
	}

	return c.Status(http.StatusOK).JSON(payload)
}

func (api *Api) NewSession(c *fiber.Ctx) error {
	req := new(NewSessionRequest)
	if err := c.BodyParser(req); err != nil {
		return api.handleError(c, e.BuildError(
			e.BAD_REQUEST,
			"can't parse the body",
		))
	}

	if err := validator.New().Struct(req); err != nil {
		return api.handleError(c, e.BuildErrorf(
			e.BAD_REQUEST,
			"validation error: %s",
			err.(validator.ValidationErrors).Error(),
		))
	}

	chess, err := game.New(api.app, req.Mode)
	if err != nil {
		return api.handleError(c, err)
	}

	resp := chess.Do(game.NewRequest().SetOperation(game.STATUS))

	return c.Status(http.StatusCreated).JSON(resp.Session)
}

func (api *Api) GetSession(c *fiber.Ctx) error {
	uuid := c.Params("id")
	if uuid == "" {
		return api.handleError(c, e.BuildError(
			e.BAD_REQUEST,
			"session id is not provided",
		))
	}

	chess, err := game.Load(api.app, uuid)
	if err != nil {
		return api.handleError(c, err)
	}

	resp := chess.Do(game.NewRequest().SetOperation(game.STATUS))

	return c.Status(http.StatusOK).JSON(resp)
}

func (api *Api) handleError(c *fiber.Ctx, err error) error {
	switch err.(type) {
	case e.NotFound:
		return c.Status(http.StatusNotFound).JSON(err.Error())
	case e.BadRequest:
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	default:
		api.app.Monitor.Notify(err)
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
}
