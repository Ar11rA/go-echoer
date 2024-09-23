// services/redis_service.go
package services

import (
	"context"
	"quote-server/utils"
)

type RedisService interface {
	SaveData(ctx context.Context, key string, value string) error
	GetData(ctx context.Context, key string) (string, error)
}

type RedisServiceImpl struct {
	RedisClient utils.RedisClient
}

func NewRedisService(redisClient utils.RedisClient) RedisService {
	return &RedisServiceImpl{RedisClient: redisClient}
}

func (r *RedisServiceImpl) SaveData(ctx context.Context, key string, value string) error {
	return r.RedisClient.Set(ctx, key, value)
}

func (r *RedisServiceImpl) GetData(ctx context.Context, key string) (string, error) {
	return r.RedisClient.Get(ctx, key)
}
