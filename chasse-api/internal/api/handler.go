package api

import (
	"chasse-api/internal/config"
	"chasse-api/internal/monitoring"
	"chasse-api/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/leonidasdeim/goconfig"
)

type ApiHandler struct {
	config  *goconfig.Config[config.Type]
	store   *store.Type
	monitor *monitoring.Type
}

func NewApiHandler(s *store.Type, c *goconfig.Config[config.Type], m *monitoring.Type) *ApiHandler {
	return &ApiHandler{
		store:   s,
		config:  c,
		monitor: m,
	}
}

func (h *ApiHandler) RegisterApiRoutes(app *fiber.App) {
	apiV1 := app.Group("/api/v1")

	health := apiV1.Group("/health")
	health.Get("", h.HealthCheck)

	session := apiV1.Group("/session")
	session.Get("/new", h.CreateSession)
	session.Get("/:sessionId", h.GetSession)
	session.Post("", h.UpdateSession)
}
