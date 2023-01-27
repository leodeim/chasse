package api

import (
	"net/http"

	e "chasse-api/internal/error"
	"chasse-api/internal/models"

	"github.com/gofiber/fiber/v2"
)

func (h *ApiHandler) HealthCheck(c *fiber.Ctx) error {
	payload := struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Config  string `json:"config"`
	}{
		h.config.GetCfg().AppName,
		h.config.GetCfg().Version,
		h.config.GetTimestamp(),
	}

	return c.Status(http.StatusOK).JSON(payload)
}

func (h *ApiHandler) CreateSession(c *fiber.Ctx) error {
	session, err := h.store.CreateSession(models.BLANK_BOARD)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.Status(http.StatusCreated).JSON(session)
}

func (h *ApiHandler) GetSession(c *fiber.Ctx) error {
	uuid := c.Params("sessionId")
	if uuid == "" {
		return h.handleError(c, e.BadRequest{Message: "sessionId not provided"})
	}

	session, err := h.store.GetSession(uuid)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(session)
}

func (h *ApiHandler) UpdateSession(c *fiber.Ctx) error {
	session := &models.SessionActionMessage{}

	if err := c.BodyParser(&session); err != nil {
		return h.handleError(c, e.BadRequest{Message: "can't parse body"})
	}

	session, err := h.store.UpdateSession(session.SessionId, session.Position)
	if err != nil {
		return h.handleError(c, err)
	}

	return c.Status(http.StatusOK).JSON(session)
}

func (h *ApiHandler) handleError(c *fiber.Ctx, err error) error {
	h.notifier.Notify(err, nil) // send error to airbrake
	switch err.(type) {
	case e.Internal:
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	case e.NotFound:
		return c.Status(http.StatusNotFound).JSON(err.Error())
	case e.BadRequest:
		return c.Status(http.StatusBadRequest).JSON(err.Error())
	default:
		return c.Status(http.StatusInternalServerError).JSON(err.Error())
	}
}
