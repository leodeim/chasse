package main

import (
	"encoding/json"
	"fmt"
	"log"

	"chasse-api/internal/api"
	"chasse-api/internal/config"
	e "chasse-api/internal/error"
	"chasse-api/internal/monitoring"
	"chasse-api/internal/socket"
	"chasse-api/internal/store"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/leonidasdeim/goconfig"
	fh "github.com/leonidasdeim/goconfig/pkg/filehandler"
)

func main() {
	log.Print("- APP START -")

	h, _ := fh.New(fh.WithName("chasse"), fh.WithType(fh.JSON))
	c, err := goconfig.Init[config.Type](h)
	if err != nil {
		log.Panicf("configuration error: %v", err)
	} else {
		print(c.GetCfg())
	}

	m := monitoring.Init(c)
	defer m.Close()
	m.Notify(e.Info{Message: "application starting"})

	app := fiber.New(fiber.Config{
		Prefork:               c.GetCfg().Prefork,
		CaseSensitive:         true,
		ServerHeader:          c.GetCfg().AppName,
		AppName:               c.GetCfg().AppName + "_" + c.GetCfg().Version,
		DisableStartupMessage: true,
	})

	s := store.Init(c)

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))
	app.Use(m.Middleware)

	socket.InitClient(app, s, m)
	api := api.NewApiHandler(s, c, m)
	api.RegisterApiRoutes(app)

	if err := app.Listen(fmt.Sprintf("%s:%s", c.GetCfg().Host, c.GetCfg().Port)); err != nil {
		log.Panic(err)
	}
}

func print(a any) {
	data, _ := json.MarshalIndent(a, "", "  ")
	log.Print(string(data))
}
