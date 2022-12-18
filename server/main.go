package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/leonidasdeim/goconfig"
	"github.com/leonidasdeim/zen-chess/server/api"
	"github.com/leonidasdeim/zen-chess/server/config"
	"github.com/leonidasdeim/zen-chess/server/socket"
	"github.com/leonidasdeim/zen-chess/server/store"
)

func main() {
	c, err := goconfig.NewConfig[config.Type](int(config.NUMBER_OF_SUBS))
	if err != nil {
		log.Panicf("Configuration error: %v", err)
	}

	app := fiber.New()
	store := store.NewStore(c.Cfg.Store.Port, c.Cfg.Store.Password)

	app.Static("/", "./assets")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	socket.SetupSocket(app, store)
	api := api.NewApiHandler(store)
	api.RegisterApiRoutes(app)

	if err := app.Listen(fmt.Sprintf("%s:%s", "localhost", c.Cfg.Port)); err != nil {
		log.Panic(err)
	}
}
