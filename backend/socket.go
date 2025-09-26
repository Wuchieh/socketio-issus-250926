package main

import (
	"context"

	"github.com/zishang520/socket.io/adapters/redis/v3"
	"github.com/zishang520/socket.io/adapters/redis/v3/adapter"
	"github.com/zishang520/socket.io/servers/socket/v3"
	"github.com/zishang520/socket.io/v3/pkg/types"
)

var socketSer *socket.Server

func initSocket() {
	options := socket.DefaultServerOptions()
	options.SetCors(&types.Cors{
		Origin:               "*",
		Methods:              "*",
		AllowedHeaders:       "*",
		Headers:              "*",
		MaxAge:               "3600",
		OptionsSuccessStatus: 209,
	})

	sRedisClient := redis.NewRedisClient(context.Background(), redisClient)

	options.SetAdapter(&adapter.RedisAdapterBuilder{
		Redis: sRedisClient,
		Opts:  adapter.DefaultRedisAdapterOptions(),
	})

	socketSer = socket.NewServer(nil, options)

	socketSer.On("connect", func(a ...any) {
		client := socketParse[*socket.Socket](a)
		client.Join("test")
	})
}

func socketParse[T any](a []any) T {
	var result T
	if len(a) > 0 {
		result = a[0].(T)
	}
	return result
}
