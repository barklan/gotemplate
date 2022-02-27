package caching

import (
	"context"
	"fmt"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-redis/redis/v8"
)

type RedisConnectionData struct {
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
}

type Redis struct {
	cl *redis.Client
}

func InitRedis() (*Redis, error) {
	cfg := RedisConnectionData{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to parse redis env vars: %w", err)
	}
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:6379", cfg.Host),
		Password: cfg.Password,
		DB:       0, // use default DB
	})
	rs := &Redis{cl: redisClient}

	if err = rs.Set("test", "", 1*time.Minute); err != nil {
		return nil, fmt.Errorf("failed to test redis: %w", err)
	}

	return rs, nil
}

func (r *Redis) Set(key string, val interface{}, ttl time.Duration) error {
	if ttl < 0 {
		ttl = 0
	}
	err := r.cl.Set(context.Background(), key, val, ttl).Err()

	return fmt.Errorf("failed to set value to redis: %w", err)
}

func (r *Redis) Get(key string) ([]byte, bool, error) {
	val, err := r.cl.Get(context.Background(), key).Result()
	switch {
	case err == redis.Nil: //nolint
		return []byte{}, false, nil
	case err != nil:
		return []byte{}, false, fmt.Errorf("failed to set key to redis: %w", err)
	default:
		return []byte(val), true, nil
	}
}
