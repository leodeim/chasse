package store

import (
	"github.com/go-redis/redis"
)

type Store struct {
	db *redis.Client
}

func NewStore() *Store {
	s := Store{}
	s.db = getNewRedisClient()
	return &s
}

func getNewRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6381",
		Password: "admin123",
		DB:       0,
	})
}
