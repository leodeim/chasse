package main

import (
	"chasse-api/core"
	e "chasse-api/error"
	"chasse-api/logger"
	"chasse-api/modules/api"
	"chasse-api/modules/socket"
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
