// utils/redis_util.go
package utils

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type RedisClient interface {
	Set(ctx context.Context, key string, value string) error
	Get(ctx context.Context, key string) (string, error)
}

type RedisUtil struct {
	Client *redis.Client
}

func (r *RedisUtil) Set(ctx context.Context, key string, value string) error {
	return r.Client.Set(ctx, key, value, 0).Err()
}

func (r *RedisUtil) Get(ctx context.Context, key string) (string, error) {
	return r.Client.Get(ctx, key).Result()
}
