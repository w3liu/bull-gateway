package main

import (
	"github.com/w3liu/bull-gateway/cmd/config"
	"github.com/w3liu/bull-gateway/pkg/server"
	"github.com/w3liu/bull-gateway/tools/log"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()
	if err := cfg.Init("./config/config.toml"); err != nil {
		panic(err)
	}
	log.Info("bull-gateway config", zap.Any("config", cfg))
	// 启动Server
	go func() {
		if err := server.Start(cfg); err != nil {
			panic(err)
		}
	}()

	select {}
}
