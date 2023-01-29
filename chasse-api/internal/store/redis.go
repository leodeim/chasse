package store

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	client     *redis.Client
	expiration time.Duration
}

func NewRedis(host string, port string, pw string, exp time.Duration) Storage {
	r := Redis{}

	r.client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pw,
		DB:       0,
	})

	return &r
}

func (r *Redis) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

func (r *Redis) Set(key string, value any) (string, error) {
	return r.client.Set(key, value, r.expiration).Result()
}
