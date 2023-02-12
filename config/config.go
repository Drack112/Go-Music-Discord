package config

import (
	"os"
)

type Config struct {
	Token string
}

var cfg *Config

func Load() (*Config, error) {
	cfg = new(Config)
	var err error

	cfg.Token = os.Getenv("TOKEN")

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
