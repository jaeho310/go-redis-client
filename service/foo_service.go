package service

import "redis-tut/gateway"

type FooService struct {
	redisGateway gateway.RedisGateway
}

func (fooService FooService) New(redisGateway gateway.RedisGateway) *FooService {
	return &FooService{redisGateway}
}

func (fooService *FooService) GetData(key string) (string, error) {
	return fooService.redisGateway.GetData(key)
}

func (fooService *FooService) SetData(key string, value string) error {
	return fooService.redisGateway.SetData(key, value)
}

func (fooService *FooService) GetKeyList() ([]string, error) {
	return fooService.redisGateway.GetKeyList()
}
