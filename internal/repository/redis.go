package repository

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"log"
)

func NewRedis(cfg *models.Redis) *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       0,
	})

	if err := client.Ping(context.TODO()).Err(); err != nil {
		log.Fatalf("an error is occured while connecting: %s", err)
	}

	return client
}
