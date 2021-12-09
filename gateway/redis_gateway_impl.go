package gateway

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type RedisGatewayImpl struct {
	client     *redis.Client
	expireTime time.Duration
}

func (redisGateway RedisGatewayImpl) New(client *redis.Client, expireTime time.Duration) *RedisGatewayImpl {
	return &RedisGatewayImpl{client: client, expireTime: expireTime}
}

func (redisGateway *RedisGatewayImpl) SetData(key string, value string) error {
	err := redisGateway.client.Set(context.Background(), key, value, redisGateway.expireTime).Err()
	if err != nil {
		return err
	}

	return nil
}

func (redisGateway *RedisGatewayImpl) GetData(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	successCh := make(chan string)
	errorCh := make(chan error)

	go func(successCh chan string, errorCh chan error) {
		result, err := redisGateway.client.Get(ctx, key).Result()
		if err != nil {
			errorCh <- err
			return
		}
		successCh <- result
	}(successCh, errorCh)

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case result := <-successCh:
		return result, nil
	case result := <-errorCh:
		return "", result
	}
}

// redis는 하나의 스레드를 갖기에 keys라는 명령어는 금기
// scan을 사용
func (redisGateway *RedisGatewayImpl) GetKeyList() ([]string, error) {
	var cursor uint64
	var keyList []string
	for {
		var keys []string
		var err error
		keys, cursor, err = redisGateway.client.Scan(context.Background(), cursor, "*", 10).Result()
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
