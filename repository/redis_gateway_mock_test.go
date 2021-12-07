package repository

import (
	"context"
	"errors"
	"github.com/go-redis/redismock/v8"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type MockRedisGatewayTestSuite struct {
	suite.Suite
	redisGateway *RedisGateway
	mock         redismock.ClientMock
	expireTime   time.Duration
}

// suite.Run()이 SetUpTest()를 실행시킵니다.
func TestMockRedisGatewayTestSuite(t *testing.T) {
	suite.Run(t, new(MockRedisGatewayTestSuite))
}

// 테스트에서 공통으로 해야하는 행위들을 구조체에 넣어줍니다.
// junit의 @beforeEach와 비슷한 기능입니다.
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) SetupTest() {
	client, mock := redismock.NewClientMock()
	redisGatewayTestSuite.expireTime = time.Second * 100
	redisGatewayTestSuite.redisGateway = RedisGateway{}.New(client, context.Background(), redisGatewayTestSuite.expireTime)
	redisGatewayTestSuite.mock = mock
}

// 생성 확인
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) TestRedisGatewayNew() {
	assert.NotNil(redisGatewayTestSuite.T(), redisGatewayTestSuite.redisGateway)
}

// redis set 테스트
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) TestRedisGatewaySet() {
	key := "hello"
	value := "world"
	redisGatewayTestSuite.mock.ExpectSet(key, value, time.Second*100).SetVal("SUCCESS")

	err := redisGatewayTestSuite.redisGateway.SetData(key, value)
	assert.NoError(redisGatewayTestSuite.T(), err)
}

// redis set 에러 테스트
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) TestRedisGatewaySetWithError() {
	key := "hello"
	value := "world"
	redisGatewayTestSuite.mock.ExpectSet(key, value, time.Second*100).SetErr(errors.New("FAIL"))
	err := redisGatewayTestSuite.redisGateway.SetData(key, value)
	assert.Error(redisGatewayTestSuite.T(), err, "FAIL")
}

// redis get 테스트
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) TestRedisGatewayGet() {
	key := "hello"
	value := "world"
	redisGatewayTestSuite.mock.ExpectGet(key).SetVal(value)
	data, err := redisGatewayTestSuite.redisGateway.GetData(key)
	assert.NoError(redisGatewayTestSuite.T(), err)
	assert.Equal(redisGatewayTestSuite.T(), value, data)
}

// redis get 에러 테스트
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) TestRedisGatewayGetWithError() {
	key := "hello"
	redisGatewayTestSuite.mock.ExpectGet(key).SetErr(errors.New("FAIL"))
	_, err := redisGatewayTestSuite.redisGateway.GetData(key)
	assert.EqualError(redisGatewayTestSuite.T(), err, "FAIL")
}

// redis scan 테스트
func (redisGatewayTestSuite *MockRedisGatewayTestSuite) TestGetKeyList() {
	var mockResult []string
	mockResult = append(mockResult, "a")
	mockResult = append(mockResult, "b")
	redisGatewayTestSuite.mock.ExpectScan(0, "*", 10).SetVal(mockResult, 0)
	list, err := redisGatewayTestSuite.redisGateway.GetKeyList()
	assert.NoError(redisGatewayTestSuite.T(), err)
	assert.EqualValues(redisGatewayTestSuite.T(), mockResult, list)

}
