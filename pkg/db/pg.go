package db

import (
	"context"
	"fmt"

	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type PGConnectionData struct {
	DB       string `env:"POSTGRES_DB"`
	Password string `env:"POSTGRES_PASSWORD"`
	User     string `env:"POSTGRES_USER"`
	Host     string `env:"POSTGRES_HOST_AND_PORT"`
}

// It's up to the caller to close connection
func Conn(lg *zap.Logger) (*pgxpool.Pool, error) {
	cfg := PGConnectionData{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("unable to parse postgres env vars: %w", err)
	}
	url := fmt.Sprintf("postgres://%s:%s@%s/%s", cfg.User, cfg.Password, cfg.Host, cfg.DB)

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to validate pg url connection string: %w", err)
	}
	config.ConnConfig.Logger = zapadapter.NewLogger(lg)

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}
	return conn, nil
}
