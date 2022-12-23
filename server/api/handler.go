package api

import (
	"github.com/leonidasdeim/goconfig"
	"github.com/leonidasdeim/zen-chess/server/config"
	"github.com/leonidasdeim/zen-chess/server/store"
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
