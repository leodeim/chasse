package api

import (
	"github.com/leonidasdeim/zen-chess/server/store"
)

type ApiHandler struct {
	store *store.Store
}

func NewApiHandler(s *store.Store) *ApiHandler {
	return &ApiHandler{
		store: s,
	}
}
