package store

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/leonidasdeim/zen-chess/server/models"
)

func (s *Store) CreateSession(position interface{}) (*models.SessionDataModel, error) {
	uuid := uuid.New().String()
	return s.UpdateSession(uuid, position)
}

func (s *Store) UpdateSession(uuid string, position interface{}) (*models.SessionDataModel, error) {
	positionString, err := json.Marshal(position)
	if err != nil {
		return nil, err
	}

	if err := s.db.Set("ses:"+uuid, positionString, 0).Err(); err != nil {
		return nil, err
	}

	return &models.SessionDataModel{
		SessionId: uuid,
		Position:  position,
	}, nil
}

func (s *Store) GetSession(uuid string) (*models.SessionDataModel, error) {
	data, err := s.db.Get("ses:" + uuid).Result()
	if err != nil {
		return nil, err
	}

	var position interface{}
	if err := json.Unmarshal([]byte(data), &position); err != nil {
		return nil, err
	}

	return &models.SessionDataModel{
		SessionId: uuid,
		Position:  position,
	}, nil
}
