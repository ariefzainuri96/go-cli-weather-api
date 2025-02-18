package main

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

func SetupRemoteRedis() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("err loading env: ", err.Error())
		return
	}

	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	rdb.Set(ctx, "foo", "bar", 0)
	result, err := rdb.Get(ctx, "foo").Result()

	if err != nil {
		panic(err)
	}

	fmt.Println(result) // >>> bar
}

func SetupLocalRedis() {
	client := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:9010",
			Password: "",
			DB:       0,
		},
	)

	ping, err := client.Ping(context.Background()).Result()

	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}

	fmt.Println("ping: ", ping)

	// try to set a key value
	err = client.Set(context.Background(), "key-from-apps", "123", 0).Err()

	if err != nil {
		fmt.Println("err: ", err.Error())
		return
	}
}
