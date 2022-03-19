package db

import (
	"context"
	"fmt"
	"time"

	retry "github.com/avast/retry-go/v4"
	"github.com/caarlos0/env"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type PGConnectionData struct {
	DB       string `env:"POSTGRES_DB"`
	Password string `env:"POSTGRES_PASSWORD"`
	User     string `env:"POSTGRES_USER"`
	Host     string `env:"POSTGRES_HOST"`
	Port     int    `env:"POSTGRES_PORT"`
}

// It's up to the caller to close connection.
func Conn(lg *zap.Logger) (*pgxpool.Pool, error) {
	cfg := PGConnectionData{}
	if err := env.Parse(&cfg); err != nil {
		return nil, fmt.Errorf("unable to parse postgres env vars: %w", err)
	}
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DB)

	config, err := pgxpool.ParseConfig(url)
	if err != nil {
		return nil, fmt.Errorf("failed to validate pg url connection string: %w", err)
	}
	config.ConnConfig.Logger = zapadapter.NewLogger(lg)

	var conn *pgxpool.Pool
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()
	if err = retry.Do(func() error {
		conn, err = pgxpool.ConnectConfig(context.Background(), config)
		return err
	}, retry.Context(ctx), retry.Delay(1*time.Second)); err != nil {
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	return conn, nil
}
