package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	API12 "gitlab.com/a5805/ondeu/ondeu-back/internal/repository/API1"
	API22 "gitlab.com/a5805/ondeu/ondeu-back/internal/repository/API2"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

type API1 interface {
	// Create creates a new love percentage
	Create(ctx context.Context, percentage models.ExternalLovePercentage) error
	// Get returns external service data from love percentage
	Get(ctx context.Context, percentage models.ExternalLovePercentage) (models.ExternalLovePercentage, error)
	// Update updates external service data from date fact in world history
	Update(ctx context.Context, in models.ExternalLovePercentage) error
}

type API2 interface {
	// Create creates a new date fact in world history
	Create(ctx context.Context, fact models.ExternalDateFact) error
	// Get returns external service data from date fact in world history
	Get(ctx context.Context, fact models.ExternalDateFact) (models.ExternalDateFact, error)
	// Update updates external service data from date fact in world history
	Update(ctx context.Context, in models.ExternalDateFact) error
}

type Repository struct {
	API2
	API1
}

func NewRepository(db *redis.Client) *Repository {
	return &Repository{
		API1: API12.NewRepository(db),
		API2: API22.NewRepository(db),
	}
}
