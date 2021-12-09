package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"redis-tut/gateway"
	"redis-tut/service"
	"time"
)

func main() {
	fooService := getFooService(getRedisGateWay())
	err := fooService.SetData("hello", "world")
	if err != nil {
		panic(err)
	}
	data, err := fooService.GetData("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	list, err := fooService.GetKeyList()
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}

func getRedisGateWay() *gateway.RedisGatewayImpl {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // password
		DB:       0,  // namespace
	})
	return gateway.RedisGatewayImpl{}.New(redisClient, time.Second*5)
}

func getFooService(redisGatewayImpl *gateway.RedisGatewayImpl) *service.FooService {
	return service.FooService{}.New(redisGatewayImpl)
}
