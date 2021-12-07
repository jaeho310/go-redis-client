package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"redis-tut/repository"
	"time"
)

func main() {
	example()
}

func example() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // password
		DB:       0,  // namespace
	})
	redisGateway := repository.RedisGateway{}.New(redisClient, context.Background(), time.Second*5)
	err := redisGateway.SetData("hello", "world")
	if err != nil {
		panic(err)
	}
	data, err := redisGateway.GetData("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
	list, err := redisGateway.GetKeyList()
	if err != nil {
		panic(err)
	}
	fmt.Println(list)
}
