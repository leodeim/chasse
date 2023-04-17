package core

import (
	"sync"

	"chasse-api/internal/logger"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberlogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

var log = logger.New("CORE")

type App struct {
	Config  *Config
	Router  *fiber.App
	Monitor *Monitor
	Store   *Store
	Sync    sync.WaitGroup
	modules map[Module]bool
}

func NewApp() *App {
	c := InitConfig()
	logger.SetGlobalLogLevel(c.Get().LogLevel)

	m := InitMonitor(c)
	s, err := InitStore(c)
	if err != nil {
		log.Error(err.Error())
		panic(err)
	}

	r := fiber.New(fiber.Config{
		Prefork:               c.Get().Prefork,
		CaseSensitive:         true,
		ServerHeader:          c.Get().AppName,
		AppName:               c.Get().Version,
		DisableStartupMessage: true,
	})

	r.Use(recover.New())
	r.Use(fiberlogger.New(fiberlogger.Config{
		TimeFormat: logger.DateTimeFormat,
		Format:     logger.FiberLogFormat,
	}))
	r.Use(cors.New())
	r.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	r.Use(m.Middleware)

	app := &App{
		Config:  c,
		Router:  r,
		Monitor: m,
		Store:   s,
		modules: make(map[Module]bool),
	}

	return app
}

func (app *App) Close() {
	log.Info("cleanup App, core components and modules")

	for mod := range app.modules {
		mod.Close()
	}
	app.Monitor.Close()
	app.Store.Close()
	app.Config.Close()
}

func (app *App) Run(mod Module) {
	app.modules[mod] = true

	app.Sync.Add(1)
	go mod.Run(func() {
		app.Sync.Done()
	})
}
