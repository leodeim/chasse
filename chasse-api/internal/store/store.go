package store

import (
	"fmt"

	"chasse-api/internal/config"

	"github.com/go-redis/redis"
	"github.com/leonidasdeim/goconfig"
)

const MODULE_NAME = "redis_store"

type Store struct {
	config *goconfig.Config[config.Type]
	db     *redis.Client
}

func NewStore(c *goconfig.Config[config.Type]) *Store {
	c.AddSubscriber(MODULE_NAME)
	s := Store{
		config: c,
	}
	s.reconfigureStore()
	go s.configurationRunner()

	return &s
}

func (s *Store) reconfigureStore() {
	config := s.config.GetCfg().Store
	s.db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0,
	})
}

func (s *Store) configurationRunner() {
	for {
		<-s.config.GetSubscriber(MODULE_NAME)
		s.reconfigureStore()
	}
}
