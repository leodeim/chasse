package store

import (
	"encoding/json"
	"fmt"
	"time"

	"chasse-api/internal/config"
	e "chasse-api/internal/error"
	"chasse-api/internal/models"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"github.com/leonidasdeim/goconfig"
)

const MODULE_NAME = "redis_store"

type Type struct {
	config *goconfig.Config[config.Type]
	db     *redis.Client
}

func Init(c *goconfig.Config[config.Type]) *Type {
	s := Type{
		config: c,
	}

	s.configure()

	s.config.AddSubscriber(MODULE_NAME)
	go s.configRunner()

	return &s
}

func (s *Type) configure() {
	config := s.config.GetCfg().Store
	s.db = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Host, config.Port),
		Password: config.Password,
		DB:       0,
	})
}

func (s *Type) configRunner() {
	for {
		<-s.config.GetSubscriber(MODULE_NAME)
		s.configure()
	}
}

func (s *Type) CreateSession(position string) (*models.SessionActionMessage, error) {
	uuid := uuid.New().String()
	return s.UpdateSession(uuid, position)
}

func (s *Type) UpdateSession(uuid string, position string) (*models.SessionActionMessage, error) {
	positionString, err := json.Marshal(position)
	if err != nil {
		return nil, e.Internal{Message: fmt.Sprintf("failed while marshal position string: %v", err)}
	}

	if err := s.db.Set("ses:"+uuid, positionString, 24*time.Hour).Err(); err != nil {
		return nil, e.Internal{Message: fmt.Sprintf("failed while writing to storage: %v", err)}
	}

	return &models.SessionActionMessage{
		SessionId: uuid,
		Position:  position,
	}, nil
}

func (s *Type) GetSession(uuid string) (*models.SessionActionMessage, error) {
	data, err := s.db.Get("ses:" + uuid).Result()
	if err != nil {
		return nil, e.NotFound{Message: fmt.Sprintf("failed while reading from storage: %v", err)}
	}

	var position string
	if err := json.Unmarshal([]byte(data), &position); err != nil {
		return nil, e.Internal{Message: fmt.Sprintf("failed while unmarshal position string: %v", err)}
	}

	return &models.SessionActionMessage{
		SessionId: uuid,
		Position:  position,
	}, nil
}
