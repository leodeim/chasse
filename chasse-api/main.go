package main

import (
	"chasse-api/internal/core"
	e "chasse-api/internal/error"
	"chasse-api/internal/logger"
	"chasse-api/internal/modules/api"
	"chasse-api/internal/modules/socket"
)

var log = logger.New("MAIN")

func main() {
	app := core.NewApp()
	defer app.Close()
	app.Monitor.Notify(e.Info{Message: "application is starting"})

	app.Run(api.New(app))
	socket.Setup(app)

	app.Sync.Wait()

	log.Info("all modules exited")
}
