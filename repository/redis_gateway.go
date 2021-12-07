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

// redis는 하나의 스레드를 갖기에 keys라는 명령어는 금기시 되어있습니다.
// scan을 이용하여 가져옵니다.
func (redisGateway *RedisGateway) GetKeyList() ([]string, error) {
	var cursor uint64
	var keyList []string
	for {
		var keys []string
		var err error
		keys, cursor, err = redisGateway.client.Scan(redisGateway.ctx, cursor, "*", 10).Result()
		if err != nil {
			return nil, err
		}
		for _, el := range keys {
			keyList = append(keyList, el)
		}
		if cursor == 0 {
			return keyList, nil
		}
	}
}
