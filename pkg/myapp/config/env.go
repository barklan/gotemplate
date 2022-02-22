package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	FastCacheSize int `env:"MYAPP_FAST_CACHE_SIZE"`
}

func Read() (*Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return &cfg, err
}
