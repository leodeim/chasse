package api

import (
	"chasse-api/internal/config"
	"chasse-api/internal/store"

	"github.com/leonidasdeim/goconfig"
)

type ApiHandler struct {
	config *goconfig.Config[config.Type]
	store  *store.Store
}

func NewApiHandler(s *store.Store, c *goconfig.Config[config.Type]) *ApiHandler {
	return &ApiHandler{
		store:  s,
		config: c,
	}
}
