package store

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/leonidasdeim/goconfig"
	"github.com/leonidasdeim/zen-chess/server/config"
)

const MODULE_NAME = "redis_store"

type Store struct {
	config *config.Connection
	notify chan bool
	db     *redis.Client
}

func NewStore(c *goconfig.Config[config.Type]) *Store {
	c.AddSubscriber(MODULE_NAME)
	s := Store{
		config: &c.GetCfg().Store,
		notify: c.GetSubscriber(MODULE_NAME),
	}
	s.reconfigureStore()
	go s.configurationRunner()

	return &s
}

func (s *Store) reconfigureStore() {
	s.db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", s.config.Host, s.config.Port),
		Password: s.config.Password,
		DB:       0,
	})
}

func (s *Store) configurationRunner() {
	for {
		<-s.notify
		s.reconfigureStore()
	}
}
