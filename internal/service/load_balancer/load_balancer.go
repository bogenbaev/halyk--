package load_balancer

import (
	"context"
	"errors"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/client"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository"
	"gitlab.com/a5805/ondeu/ondeu-back/internal/repository/cache_redis"
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

	if clientData, err := s.repo.Get(ctx, in); !errors.Is(err, cache_redis.ErrNotFound) && err != nil {
		return resp, err
	} else if errors.Is(err, cache_redis.ErrNotFound) {
		return resp, s.repo.Cache.Set(ctx, in)
	} else {
		return clientData, nil
	}
}

func (s *service) getRandomURLByWeight() string {
	totalWeight := 0
	for _, balance := range s.cfg.Balances {
		totalWeight += balance.Weight
	}

	rValue := rand.Intn(totalWeight)
	left, right := 0, 0
	for _, balance := range s.cfg.Balances {
		left = right
		right += balance.Weight

		if rValue > left && rValue <= right {
			return balance.Url
		}
	}

	return ""
}
