package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	Secret string `env:"MYAPP_SECRET"`
}

func Read() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("failed to parse config from env vars: %w", err)
	}

	return &cfg, nil
}
