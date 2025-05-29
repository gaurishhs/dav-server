package config

import (
	"github.com/BurntSushi/toml"
)

type ServerConfig struct {
	Addr string `toml:"addr"`
}

type Config struct {
	Server ServerConfig `toml:"server"`
}

func LoadConfig(configPath string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		return nil, err
	}
	return &config, nil
}
