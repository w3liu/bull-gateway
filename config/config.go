package config

import "github.com/BurntSushi/toml"

type Config struct {
	Runtime    string
	ServerPort string
	ManagePort string
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
