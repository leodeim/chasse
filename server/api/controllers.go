package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leonidasdeim/zen-chess/server/models"
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
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(session)
}

func (h *ApiHandler) GetSession(c *fiber.Ctx) error {
	uuid := c.Params("sessionId")
	if uuid == "" {
		return c.Status(http.StatusNotFound).JSON("Wrong Session ID")
	}

	session, err := h.store.GetSession(uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(session)
}

func (h *ApiHandler) UpdateSession(c *fiber.Ctx) error {
	session := &models.SessionActionMessage{}

	if err := c.BodyParser(&session); err != nil {
		return err
	}

	session, err := h.store.UpdateSession(session.SessionId, session.Position)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(session)
}
