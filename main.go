package main

import (
	"github.com/w3liu/bull-gateway/config"
	"github.com/w3liu/bull-gateway/pkg/log"
	"github.com/w3liu/bull-gateway/server"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()
	if err := cfg.Init("./config/config.toml"); err != nil {
		panic(err)
	}

	// 启动Server
	go func() {
		if err := server.Run(cfg); err != nil {
			panic(err)
		}
	}()

	log.Info("bull-gateway start success", zap.Any("config", cfg))

	select {}
}
