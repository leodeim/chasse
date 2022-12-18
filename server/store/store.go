package store

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/leonidasdeim/zen-chess/server/config"
)

type Store struct {
	config *config.Connection
	notify chan bool
	db     *redis.Client
}

func NewStore(c *config.Connection, ch chan bool) *Store {
	s := Store{
		config: c,
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
