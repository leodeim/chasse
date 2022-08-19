package api

import (
	"github.com/gofiber/fiber/v2"
)

func (h *ApiHandler) RegisterApiRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	health := apiV1.Group("/health")
	health.Get("", h.HealthCheck)

	session := apiV1.Group("/session")
	session.Get("/new", h.CreateSession)
	session.Get("/:uuid", h.GetSession)
	session.Post("", h.UpdateSession)
}
