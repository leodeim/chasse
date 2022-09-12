package socket

import (
	"github.com/gofiber/fiber/v2"
	"github.com/leonidasdeim/zen-chess/server/store"
)

func SetupSocket(app *fiber.App, store *store.Store) {
	serveClient(app, store)
}
