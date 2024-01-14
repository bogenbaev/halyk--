package service

import (
	"context"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/client"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/service/load_balancer"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Proxifier interface {
	// Proxify implementing load balancing simulation...
	Proxify(context.Context, models.Request) (models.Response, error)
}

type Services struct {
	Proxifier
}

func NewServices(cfg *models.AppConfigs, info *client.ExternalServices, repo *repository.Repository) *Services {
	return &Services{
		Proxifier: load_balancer.NewLoadBalancerService(cfg, info, repo),
	}
}
