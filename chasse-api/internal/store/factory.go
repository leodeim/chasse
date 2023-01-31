package store

import (
	"chasse-api/internal/config"
	"chasse-api/internal/store/impl"
	"strings"
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

func StorageFactory(c config.Storage) Storage {
	switch strings.ToUpper(c.Type) {
	case REDIS:
		return impl.NewRedis(c.Host, c.Port, c.Password, c.Expiration)
	case BADGER:
		return impl.NewBadgerDB(c.InMemory, c.FileLocation)
	default:
		return impl.NewRedis(c.Host, c.Port, c.Password, c.Expiration)
	}
}
