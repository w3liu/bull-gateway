package server

import (
	"github.com/w3liu/bull-gateway/config"
	"github.com/w3liu/bull-gateway/server/handler"
	"net/http"
)

func Run(cfg *config.Config) error {
	return http.ListenAndServe(cfg.ServerPort, handler.New())
}
