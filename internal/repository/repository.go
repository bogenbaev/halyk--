package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository/cache_redis"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

type Cache interface {
	// Set creates a new love percentage
	Set(ctx context.Context, percentage models.Request) error
	// Get returns client service data from love percentage
	Get(ctx context.Context, percentage models.Request) (models.Response, error)
}

type Repository struct {
	Cache
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{
		Cache: cache_redis.NewCache(db),
	}
}
