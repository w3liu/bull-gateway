package api

import (
	"github.com/w3liu/bull-gateway/pkg/settings/internal/service"
	"github.com/w3liu/bull-gateway/pkg/settings/internal/store"
)

type Handler struct {
	srv service.Service
}

func newHandler(s store.Factory) *Handler {
	return &Handler{srv: service.NewService(s)}
}
