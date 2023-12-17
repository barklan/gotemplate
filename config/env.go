package config

import (
	"fmt"

	"github.com/caarlos0/env/v6"
)

const prefix = "MYAPP_"

type Config struct {
	Secret string `env:"SECRET" envDefault:"12345"`
}

func Read() (*Config, error) {
	cfg := Config{}
	if err := env.Parse(&cfg, env.Options{
		Prefix: prefix,
	}); err != nil {
		return nil, fmt.Errorf("failed to parse config from env vars: %w", err)
	}

	return &cfg, nil
}

