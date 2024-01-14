package service

import (
	"context"
	"gitlab.com/a5805/ondeu/ondeu-back/external"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	API12 "gitlab.com/a5805/ondeu/ondeu-back/internal/service/API1"
	API22 "gitlab.com/a5805/ondeu/ondeu-back/internal/service/API2"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/service/calculate_percent_division"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type CalculatePercentDivision interface {
	// CalculatePercentDivision calculate separate logic
	CalculatePercentDivision(context.Context) (*models.PercentageDivision, error)
}

type API1 interface {
	// Create creates a new external service1 data
	Create(ctx context.Context, percentage models.ExternalLovePercentage) error
	// Get returns data from redis
	Get(ctx context.Context, percentage models.ExternalLovePercentage) (models.ExternalLovePercentage, error)
	// Update updates data in redis
	Update(ctx context.Context, percentage models.ExternalLovePercentage) error
}

type API2 interface {
	// Create creates a new external service2 data
	Create(ctx context.Context, percentage models.ExternalDateFact) error
	// Get returns data from redis
	Get(ctx context.Context, percentage models.ExternalDateFact) (models.ExternalDateFact, error)
	// Update updates data in redis
	Update(ctx context.Context, percentage models.ExternalDateFact) error
}

type Services struct {
	CalculatePercentDivision
	API1
	API2
}

func NewServices(cfg *models.AppConfigs, info *external.ExternalServices, repo *repository.Repository) *Services {
	return &Services{
		CalculatePercentDivision: calculate_percent_division.NewPercentDivisionService(cfg, info, repo),
		API1:                     API12.NewAPI1Service(repo.API1),
		API2:                     API22.NewAPI2Service(repo.API2),
	}
}
