package cache_redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
	"time"
)

var ErrNotFound = errors.New("no matching record found in redis database")

type cache struct {
	db *redis.Client
}

func NewCache(db *redis.Client) *cache {
	return &cache{
		db: db,
	}
}

func (r *cache) key(in models.Request) string {
	return fmt.Sprintf("url:%s:%s", in.Url, in.IP)
}

func (r *cache) Set(ctx context.Context, in models.Request) error {
	data, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to encode number: %w", err)
	}

	// Получаем текущее время
	now := time.Now()

	// Вычисляем длительность до полуночи следующего дня
	durationUntilNextDay := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location()).Sub(now)

	key := r.key(in)

	txn := r.db.TxPipeline()
	res := txn.SetNX(ctx, key, string(data), durationUntilNextDay)
	if err = res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to set: %w", err)
	}

	if err = txn.SAdd(ctx, "loves", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add to client set: %w", err)
	}

	if _, err = txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (r *cache) Get(ctx context.Context, in models.Request) (models.Response, error) {
	key := r.key(in)

	value, err := r.db.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return models.Response{}, ErrNotFound
	} else if err != nil {
		return models.Response{}, fmt.Errorf("get order: %w", err)
	}

	var res models.Response
	err = json.Unmarshal([]byte(value), &res)
	if err != nil {
		return models.Response{}, fmt.Errorf("failed to decode client json: %w", err)
	}

	return res, nil
}
