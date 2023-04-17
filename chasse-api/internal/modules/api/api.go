package api

import (
	"chasse-api/internal/core"
	"chasse-api/internal/logger"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2/middleware/monitor"
)

var log = logger.New("API")

type Api struct {
	app *core.App
}

func New(app *core.App) *Api {
	api := &Api{
		app: app,
	}

	api.registerRoutes()

	return api
}

func (api *Api) registerRoutes() {
	router := api.app.Router
	v1 := router.Group("/api/v1")

	v1.Post("/session", api.NewSession)
	v1.Get("/session/:id", api.GetSession)

	v1.Get("/metrics", monitor.New(monitor.Config{Title: "metrics"}))
	v1.Get("/health", api.HealthCheck)
}

func (api *Api) Run(cleanup func()) {
	defer cleanup()

	r := api.app.Router
	c := api.app.Config
	addr := fmt.Sprintf("%s:%s", c.Get().Host, c.Get().Port)

	log.Infof("server started on: %s", addr)
	go func() {
		if err := r.Listen(addr); err != nil {
			log.Error(err.Error())
			panic(err)
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt
}

func (api *Api) Close() {
	log.Info("closing API module")

	go func() {
		router := api.app.Router
		err := router.Shutdown()

		if err != nil {
			log.Error(err.Error())
			panic(err)
		}
	}()
}
