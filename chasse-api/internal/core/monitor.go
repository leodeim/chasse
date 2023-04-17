package core

import (
	airbrake "github.com/airbrake/gobrake/v5"
	fiberbrake "github.com/airbrake/gobrake/v5/fiber"
	"github.com/gofiber/fiber/v2"
)

var SKIP_PATHS = map[string]struct{}{
	"/api/v1/health":  {},
	"/api/v1/metrics": {},
	"/api/ws":         {},
}

type Monitor struct {
	notifier *airbrake.Notifier
	handler  fiber.Handler
	status   bool
}

func InitMonitor(c *Config) *Monitor {
	m := Monitor{
		notifier: nil,
		handler:  nil,
	}

	m.configure(c.Get())
	c.config.AddCallback(
		func(c Configuration) {
			m.configure(c)
		},
	)

	return &m
}

func (m *Monitor) Notify(err error) {
	if m.status {
		m.notifier.Notify(err, nil)
	}
}

func (m *Monitor) Close() {
	if m.status {
		log.Info("monitor: closing")
		m.notifier.Close()
	}
	m.status = false
}

func (m *Monitor) Middleware(c *fiber.Ctx) error {
	path := c.Request().URI().Path()

	if _, ok := SKIP_PATHS[string(path)]; !ok {
		if m.status && m.handler != nil {
			return m.handler(c)
		}
	}

	return c.Next()
}

func (m *Monitor) configure(config Configuration) error {
	log.Info("monitor: configure")

	if !m.isConfigAvailable(config) {
		m.Close()
		return nil
	}

	m.notifier = airbrake.NewNotifierWithOptions(&airbrake.NotifierOptions{
		ProjectId:   config.Monitor.Id,
		ProjectKey:  config.Monitor.Key,
		Environment: config.Monitor.Env,
	})
	m.handler = fiberbrake.New(m.notifier)
	m.status = true

	return nil
}

func (m *Monitor) isConfigAvailable(config Configuration) bool {
	return config.Monitor.Key != "undefined" &&
		config.Monitor.Id != 0
}
