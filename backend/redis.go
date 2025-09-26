package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

var redisClient *redis.Client

func initRedis() error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "redis_y3KG57",
		DB:       1,
	})

	return redisClient.Ping(context.Background()).Err()
}
