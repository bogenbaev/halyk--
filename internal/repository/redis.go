package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func NewRedis(address string) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		log.Fatalf("an error is occured while connecting: %s", err)
	}

	return client
}
