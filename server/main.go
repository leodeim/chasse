package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/leonidasdeim/zen-chess/server/api"
	"github.com/leonidasdeim/zen-chess/server/socket"
	"github.com/leonidasdeim/zen-chess/server/store"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()
	store := store.NewStore(os.Getenv("REDIS_PORT"), os.Getenv("REDIS_PW"))

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

	if err := app.Listen("localhost:8085"); err != nil {
		log.Panic(err)
	}
}
