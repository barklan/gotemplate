package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	MyAppEnv string `env:"MY_APP_ENV"`
}

func Read() (Config, error) {
	cfg := Config{}
	err := env.Parse(&cfg)
	return cfg, err
}
