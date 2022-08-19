package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/leonidasdeim/zen-chess/internal/api"
	"github.com/leonidasdeim/zen-chess/internal/socket"
	"github.com/leonidasdeim/zen-chess/internal/store"
)

func main() {
	app := fiber.New()
	store := store.NewStore()

	app.Static("/", "./assets")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	socket.Setup(app, store)
	api := api.NewApiHandler(store)
	api.RegisterApiRoutes(app)

	if err := app.Listen("localhost:8085"); err != nil {
		log.Panic(err)
	}
}
