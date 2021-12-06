package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisGateway struct {
	client     *redis.Client
	ctx        context.Context
	expireTime time.Duration
}

func (redisGateway RedisGateway) New(client *redis.Client, ctx context.Context, expireTime time.Duration) *RedisGateway {
	return &RedisGateway{client: client, ctx: ctx, expireTime: expireTime}
}

func (redisGateway *RedisGateway) SetData(key string, value string) error {
	err := redisGateway.client.Set(redisGateway.ctx, key, value, redisGateway.expireTime).Err()
	if err != nil {
		return err
	}
	return nil
}

func (redisGateway *RedisGateway) GetData(key string) (string, error) {
	result, err := redisGateway.client.Get(redisGateway.ctx, key).Result()
	if err != nil {
		return "", err
	}
	return result, nil
}
