package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func ConnectToRedis() (*redis.Client, error) {
	redisAddr := os.Getenv("REDIS_ADDR")
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisDB := os.Getenv("REDIS_DB")

	dbIndex := 0
	if redisDB != "" {
		dbIndex, _ = strconv.Atoi(redisDB)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       dbIndex,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("unable to connect to Redis: %v", err)
	}

	return rdb, nil
}

func CloseRedis() {
	if rdb != nil {
		err := rdb.Close()
		if err != nil {
			log.Printf("Error closing Redis connection: %v", err)
		}
	}
}
