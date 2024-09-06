package cache

import (
	"context"
	"time"

	"github.com/Ndraaa15/ConnectMe/internal/adapter/pkg/errx"
	"github.com/Ndraaa15/ConnectMe/internal/core/port"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
)

type Redis struct {
	redis *redis.Client
}

func NewRedisClient(redis *redis.Client) port.CacheItf {
	return &Redis{redis: redis}
}

func (r *Redis) Set(ctx context.Context, key string, value interface{}, exp time.Duration) error {
	err := r.redis.Set(ctx, key, value, exp).Err()
	if err != nil {
		return errx.New(fiber.StatusInternalServerError, "Failed to set key", err)
	}

	return nil
}

func (r *Redis) Get(ctx context.Context, key string) (string, error) {
	val, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		return "", errx.New(fiber.StatusInternalServerError, "Failed to get key", err)
	}

	return val, nil
}

func (r *Redis) Delete(ctx context.Context, key string) error {
	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return errx.New(fiber.StatusInternalServerError, "Failed to delete key", err)
	}

	return nil
}
