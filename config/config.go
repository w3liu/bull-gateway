package config

import (
	"github.com/BurntSushi/toml"
	"github.com/w3liu/bull-gateway/infra/mysql"
)

type Config struct {
	Runtime    string
	ServerAddr string
	ManageAddr string
	Service    Service
	Mysql      mysql.Conf
}

type Service struct {
	Name     string
	Registry []string
}

func New() *Config {
	return &Config{}
}

func (c *Config) Init(path string) error {
	_, err := toml.DecodeFile(path, c)
	return err
}
