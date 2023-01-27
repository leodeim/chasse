package monitoring

import (
	"chasse-api/internal/config"

	"github.com/airbrake/gobrake/v5"
	fiberbrake "github.com/airbrake/gobrake/v5/fiber"
	"github.com/gofiber/fiber/v2"
	"github.com/leonidasdeim/goconfig"
)

const MODULE_NAME = "monitoring"

var MAINTENANCE_PATHS = map[string]struct{}{
	"/api/v1/health":  {},
	"/api/v1/metrics": {},
}

type Type struct {
	config   *goconfig.Config[config.Type]
	notifier *gobrake.Notifier
	handler  fiber.Handler
	status   bool
}

func Init(c *goconfig.Config[config.Type]) *Type {
	m := Type{
		config:   c,
		notifier: nil,
		handler:  nil,
	}

	m.configure()

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

func (m *Type) Middleware(c *fiber.Ctx) error {
	path := c.Request().URI().Path()

	if _, ok := MAINTENANCE_PATHS[string(path)]; !ok {
		if m.status && m.handler != nil {
			return m.handler(c)
		}
	}

	return c.Next()
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
	m.handler = fiberbrake.New(m.notifier)
	m.status = true
}

func (m *Type) configRunner() {
	for {
		<-m.config.GetSubscriber(MODULE_NAME)
		m.configure()
	}
}

func (m *Type) isConfigAvailable() bool {
	return m.config.GetCfg().Monitoring.Key != "undefined" &&
		m.config.GetCfg().Monitoring.Id != 0
}
