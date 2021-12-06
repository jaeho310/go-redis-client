package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"redis-tut/repository"
	"time"
)

func main() {
	ExampleClient()
}

func ExampleClient() {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	redisGateway := repository.RedisGateway{}.New(redisClient, context.Background(), time.Second*100)
	err := redisGateway.SetData("hello", "world")
	if err != nil {
		panic(err)
	}
	data, err := redisGateway.GetData("hello")
	if err != nil {
		panic(err)
	}
	fmt.Println(data)
}
