package storages

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	client     *redis.Client
	expiration time.Duration
	status     bool
}

func NewRedis(host string, port string, pw string, exp int) (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: pw,
		DB:       0,
	})

	return &Redis{
		client:     client,
		expiration: time.Duration(exp) * time.Hour,
		status:     true,
	}, nil
}

func (r *Redis) Get(key string) ([]byte, error) {
	v, e := r.client.Get(key).Result()
	return []byte(v), e
}

func (r *Redis) Set(key string, value []byte) error {
	_, err := r.client.Set(key, value, r.expiration).Result()
	return err
}

func (r *Redis) Close() error {
	r.status = false
	return r.client.Close()
}

func (r *Redis) Status() bool {
	return r.status
}
