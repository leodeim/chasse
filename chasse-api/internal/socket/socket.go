package socket

import (
	"chasse-api/internal/store"

	"github.com/gofiber/fiber/v2"
)

func SetupSocket(app *fiber.App, store *store.Store) {
	serveClient(app, store)
}
