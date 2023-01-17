package main

import (
	"fmt"
	"log"

	"chasse-api/internal/api"
	"chasse-api/internal/config"
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
	h, _ := fh.New(fh.WithName("chasse"), fh.WithType(fh.JSON))
	c, err := goconfig.Init[config.Type](h)
	if err != nil {
		log.Panicf("Configuration error: %v", err)
	}

	app := fiber.New()
	store := store.NewStore(c)

	app.Static("/", "./assets")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	socket.SetupSocket(app, store)
	api := api.NewApiHandler(store, c)
	api.RegisterApiRoutes(app)

	if err := app.Listen(fmt.Sprintf("%s:%s", "localhost", c.GetCfg().Port)); err != nil {
		log.Panic(err)
	}
}
