package monitoring

import (
	"chasse-api/internal/config"

	"github.com/airbrake/gobrake/v5"
	fiberbrake "github.com/airbrake/gobrake/v5/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/leonidasdeim/goconfig"
)

const MODULE_NAME = "monitoring"

type Type struct {
	config   *goconfig.Config[config.Type]
	notifier *gobrake.Notifier
	status   bool
}

func Init(app *fiber.App, c *goconfig.Config[config.Type]) *Type {
	m := Type{
		config:   c,
		notifier: nil,
	}

	app.Use(fiberbrake.New(m.notifier))

	m.config.AddSubscriber(MODULE_NAME)
	go m.configRunner()

	return &m
}

func (m *Type) Notify(err error) {
	if m.status {
		m.notifier.Notify(err, nil)
	}
}

func (m *Type) Close() {
	if m.status {
		m.notifier.Close()
	}
	m.status = false
}

func (m *Type) configure() {
	if !m.isConfigAvailable() {
		m.Close()
		return
	}

	m.notifier = gobrake.NewNotifierWithOptions(&gobrake.NotifierOptions{
		ProjectId:   m.config.GetCfg().Monitoring.Id,
		ProjectKey:  m.config.GetCfg().Monitoring.Key,
		Environment: m.config.GetCfg().Monitoring.Environment,
	})

	m.status = true
}

func (m *Type) configRunner() {
	for {
		m.configure()
		<-m.config.GetSubscriber(MODULE_NAME)
	}
}

func (m *Type) isConfigAvailable() bool {
	return m.config.GetCfg().Monitoring.Key != "undefined" &&
		m.config.GetCfg().Monitoring.Id != 0
}
