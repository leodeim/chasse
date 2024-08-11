package socket

import (
	"chasse-api/core"
	"chasse-api/logger"
)

var log = logger.New("WS")

const path = "/api/ws"

func Setup(app *core.App) {
	log.Infof("server registered on path: %s", path)
	router := app.Router
	router.Get(path, handleClient(app))
}
