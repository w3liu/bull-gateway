package server

import (
	"github.com/w3liu/bull-gateway/config"
	"github.com/w3liu/bull-gateway/server/handler"
	"github.com/w3liu/bull/web"
)

func Start(cfg *config.Config) error {
	var service web.Service
	service = web.NewService(
		web.Name(cfg.Service.Name),
		web.Address(cfg.ServerAddr),
	)
	service.Handle("/", handler.New())
	service.Init()
	err := service.Run()
	return err
}
