package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"encoding/json"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

type RedisService struct {
	Client *redis.Client
}

func NewRedisService() *RedisService {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("err loading env:", err.Error())
		return nil
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Username: os.Getenv("REDIS_USERNAME"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       0,
	})

	return &RedisService{Client: client}
}

// func (r *RedisService) SetValue(key, value string) error {
// 	ctx := context.Background()
// 	return r.Client.Set(ctx, key, value, 0).Err()
// }

// func (r *RedisService) GetValue(key string) (string, error) {
// 	ctx := context.Background()
// 	return r.Client.Get(ctx, key).Result()
// }

// Generic function to set any type of value in Redis
func SetRedisValue[T any](rdb *redis.Client, key string, value T, expiration time.Duration) error {
	ctx := context.Background()

	// Convert value to JSON string
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}

	return rdb.Set(ctx, key, data, expiration).Err()
}

// Generic function to get a value from Redis and decode it into the desired type
func GetRedisValue[T any](rdb *redis.Client, key string) (T, error) {
	ctx := context.Background()

	var result T

	data, err := rdb.Get(ctx, key).Result()
	if err != nil {
		return result, err
	}

	// Decode JSON back to the requested type
	err = json.Unmarshal([]byte(data), &result)
	return result, err
}

func SetupRemoteRedis() {
	// err := godotenv.Load()

	// if err != nil {
	// 	fmt.Println("err loading env: ", err.Error())
	// 	return
	// }

	// ctx := context.Background()

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     os.Getenv("REDIS_ADDR"),
	// 	Username: os.Getenv("REDIS_USERNAME"),
	// 	Password: os.Getenv("REDIS_PASS"),
	// 	DB:       0,
	// })

	// rdb.Set(ctx, "foo", "bar", 0)
	// result, err := rdb.Get(ctx, "foo").Result()

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(result) // >>> bar
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
