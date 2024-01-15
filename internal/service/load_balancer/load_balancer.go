package load_balancer

import (
	"context"
	"errors"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/client"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository/cache"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"math/rand"
)

type service struct {
	cfg  *models.AppConfigs
	info *client.ExternalServices
	repo *repository.Repository
}

func NewLoadBalancerService(cfg *models.AppConfigs, info *client.ExternalServices, repo *repository.Repository) *service {
	return &service{
		info: info,
		cfg:  cfg,
		repo: repo,
	}
}

func (s *service) Proxify(ctx context.Context, in models.Request) (models.Response, error) {
	url := s.getRandomURLByWeight()

	resp, err := s.info.Proxifier.SendRequest(ctx, url)
	if err != nil {
		return resp, err
	}

	if clientData, err := s.repo.Get(ctx, in); !errors.Is(err, cache.ErrNotFound) && err != nil {
		return resp, err
	} else if errors.Is(err, cache.ErrNotFound) {
		return resp, s.repo.Cache.Set(ctx, in)
	} else {
		return clientData, nil
	}
}

func (s *service) getRandomURLByWeight() string {
	type Balance struct {
		url    string
		weight float64
	}

	balances := []Balance{
		{
			url:    s.cfg.API1,
			weight: s.cfg.Ratio.Api1Percent,
		},
		{
			url:    s.cfg.API2,
			weight: s.cfg.Ratio.Api2Percent,
		},
	}

	totalWeight := 0.0
	for _, balance := range balances {
		totalWeight += balance.weight
	}

	rValue := rand.Intn(int(totalWeight * 100))
	left, right := 0.0, 0.0
	for _, balance := range balances {
		left = right
		right += balance.weight * 100

		if rValue > int(left) && rValue <= int(right) {
			return balance.url
		}
	}

	return ""
}
