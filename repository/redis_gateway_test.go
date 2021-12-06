package repository

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type RedisGatewayTestSuite struct {
	suite.Suite
	redisGateway *RedisGateway
}

// suite.Run()이 SetUpTest()를 실행시킵니다.
func TestRedisGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(RedisGatewayTestSuite))
}

// 테스트에서 공통으로 해야하는 행위들을 구조체에 넣어줍니다.
// junit의 @beforeEach와 비슷한 기능입니다.
func (redisGatewayTestSuite *RedisGatewayTestSuite) SetupTest() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisGatewayTestSuite.redisGateway = RedisGateway{}.New(redisClient, context.Background(), time.Second*100)
}

// 생성 확인
func (redisGatewayTestSuite *RedisGatewayTestSuite) TestRedisGatewayNew() {
	assert.NotNil(redisGatewayTestSuite.T(), redisGatewayTestSuite.redisGateway)
}

func (redisGatewayTestSuite *RedisGatewayTestSuite) TestRedisGatewaySetAndGet() {
	key := "hello"
	value := "world"
	err := redisGatewayTestSuite.redisGateway.SetData(key, value)
	assert.NoError(redisGatewayTestSuite.T(), err)
	data, err := redisGatewayTestSuite.redisGateway.GetData(key)
	assert.NoError(redisGatewayTestSuite.T(), err)
	assert.Equal(redisGatewayTestSuite.T(), value, data)
}