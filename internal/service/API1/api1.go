package API1

import (
	"context"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

type service struct {
	repo repository.API1
}

func NewAPI1Service(repo repository.API1) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, percentage models.ExternalLovePercentage) error {
	return s.repo.Create(ctx, percentage)
}

func (s *service) Get(ctx context.Context, percentage models.ExternalLovePercentage) (models.ExternalLovePercentage, error) {
	return s.repo.Get(ctx, percentage)
}

func (s *service) Update(ctx context.Context, percentage models.ExternalLovePercentage) error {
	return s.repo.Update(ctx, percentage)
}
