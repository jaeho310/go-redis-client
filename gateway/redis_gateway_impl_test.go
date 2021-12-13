package gateway

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"strconv"
	"testing"
	"time"
)

// 실제로 테스트할 구현체와 suite.Suite를 담는 struct를 만들어줍니다.
type RedisGatewayTestSuite struct {
	suite.Suite
	redisGateway *RedisGatewayImpl
}

// suite.Run()이 SetUpTest()를 실행시킵니다.
func TestRedisGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(RedisGatewayTestSuite))
}

// 모든 함수마다 RedisGatewayImpl을 구축할 수는 없습니다.
// 테스트에서 공통으로 해야하는 행위들을 넣어줍니다.
// junit의 @beforeEach와 비슷한 기능입니다.
func (redisGatewayTestSuite *RedisGatewayTestSuite) SetupTest() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // password
		DB:       0,  // namespace
	})
	ctx := context.Background()
	redisClient.FlushAll(ctx) // 실제 상태를 테스트하므로 flush를 해줍니다.
	redisGatewayTestSuite.redisGateway = RedisGatewayImpl{}.New(redisClient, time.Second*100)
}

// 생성 확인
func (redisGatewayTestSuite *RedisGatewayTestSuite) TestRedisGatewayNew() {
	assert.NotNil(redisGatewayTestSuite.T(), redisGatewayTestSuite.redisGateway)
}

// set과 get 확인
func (redisGatewayTestSuite *RedisGatewayTestSuite) TestRedisGatewaySetAndGet() {
	key := "hello"
	value := "world"
	err := redisGatewayTestSuite.redisGateway.SetData(key, value)
	assert.NoError(redisGatewayTestSuite.T(), err)
	data, err := redisGatewayTestSuite.redisGateway.GetData(key)
	assert.NoError(redisGatewayTestSuite.T(), err)
	assert.Equal(redisGatewayTestSuite.T(), value, data)
}

// key list 테스트
func (redisGatewayTestSuite *RedisGatewayTestSuite) TestRedisGatewayGetKeyList() {
	cnt := 15
	for i := 0; i < cnt; i++ {
		err := redisGatewayTestSuite.redisGateway.SetData(strconv.Itoa(i), strconv.Itoa(i))
		assert.NoError(redisGatewayTestSuite.T(), err)
	}
	res, err := redisGatewayTestSuite.redisGateway.GetKeyList()
	assert.NoError(redisGatewayTestSuite.T(), err)
	assert.Equal(redisGatewayTestSuite.T(), cnt, len(res))
}
