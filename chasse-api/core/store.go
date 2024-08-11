package core

import (
	"strings"
	"sync"

	e "chasse-api/error"
	"chasse-api/storages"
)

type Storage interface {
	Get(string) ([]byte, error)
	Set(string, []byte) error
	Close() error
	Status() bool
}

const (
	BADGER string = "BADGER"
	REDIS  string = "REDIS"
)

func BuildStorage(c StoreConfig) (Storage, error) {
	switch strings.ToUpper(c.Type) {
	case REDIS:
		return storages.NewRedis(c.Host, c.Port, c.Password, c.Expiration)
	case BADGER:
		return storages.NewBadgerDB(c.InMemory, c.Location)
	default:
		return storages.NewRedis(c.Host, c.Port, c.Password, c.Expiration)
	}
}

type Store struct {
	mutex sync.Mutex
	db    Storage
}

func InitStore(c *Config) (*Store, error) {
	s := Store{}

	err := s.configure(c.Get())
	if err != nil {
		return nil, err
	}

	c.cog.AddCallback(
		func(c Configuration) {
			s.configure(c)
		},
	)

	return &s, nil
}

func (s *Store) Close() {
	if s.db != nil && s.db.Status() {
		log.Info("store: closing")
		s.db.Close()
	}
}

func (s *Store) configure(config Configuration) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	log.Info("store: configure")

	if s.db != nil && s.db.Status() {
		s.db.Close()
	}

	var err error = nil
	s.db, err = BuildStorage(config.Store)

	return err
}

func (s *Store) Write(id string, data []byte) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if err := s.db.Set("ses:"+id, data); err != nil {
		return e.BuildErrorf(e.INTERNAL, "failed while writing to storage: %s", err.Error())
	}

	return nil
}

func (s *Store) Read(id string) ([]byte, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := s.db.Get("ses:" + id)
	if err != nil {
		return nil, e.BuildErrorf(e.NOT_FOUND, "failed while reading from storage: %s", err.Error())
	}

	return data, nil
}
