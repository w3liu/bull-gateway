package main

import (
	"flag"
	"github.com/w3liu/bull-gateway/config"
	"github.com/w3liu/bull-gateway/infra/log"
	"github.com/w3liu/bull-gateway/pkg/server"
	"go.uber.org/zap"
)

var configPath = flag.String("c", "config.toml", "配置文件路径")

func main() {
	flag.Parse()
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
