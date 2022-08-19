package api

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/leonidasdeim/zen-chess/internal/models"
)

func (h *ApiHandler) HealthCheck(c *fiber.Ctx) error {
	payload := struct {
		Name    string `json:"name"`
		Version string `json:"version"`
		Config  string `json:"config"`
	}{
		"AppName",
		"Version",
		"Timestamp",
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
	uuid := c.Params("uuid")
	if uuid == "" {
		return c.Status(http.StatusNotFound).JSON("Wrong UUID")
	}

	session, err := h.store.GetSession(uuid)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(session)
}

func (h *ApiHandler) UpdateSession(c *fiber.Ctx) error {
	session := &models.SessionDataModel{}

	if err := c.BodyParser(&session); err != nil {
		return err
	}

	session, err := h.store.UpdateSession(session.Uuid, session.Position)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(err)
	}

	return c.Status(http.StatusOK).JSON(session)
}
