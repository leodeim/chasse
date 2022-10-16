package store

import (
	"fmt"

	"github.com/go-redis/redis"
)

type Store struct {
	db *redis.Client
}

func NewStore(port string, pw string) *Store {
	s := Store{}
	s.db = getNewRedisClient(port, pw)
	return &s
}

func getNewRedisClient(port string, pw string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("localhost:%s", port),
		Password: pw,
		DB:       0,
	})
}
