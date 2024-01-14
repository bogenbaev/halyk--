package API1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gitlab.com/a5805/ondeu/ondeu-back/pkg/models"
)

var ErrNotFound = errors.New("no matching record found in redis database")

type repository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) *repository {
	return &repository{
		db: db,
	}
}

func (r *repository) key(in models.ExternalLovePercentage) string {
	return fmt.Sprintf("user:%s:%s", in.FName, in.FName)
}

func (r *repository) Create(ctx context.Context, in models.ExternalLovePercentage) error {
	data, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to encode number: %w", err)
	}

	key := r.key(in)

	txn := r.db.TxPipeline()
	res := txn.SetNX(ctx, key, string(data), 0)
	if err = res.Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to set: %w", err)
	}

	if err = txn.SAdd(ctx, "loves", key).Err(); err != nil {
		txn.Discard()
		return fmt.Errorf("failed to add to numbers set: %w", err)
	}

	if _, err = txn.Exec(ctx); err != nil {
		return fmt.Errorf("failed to exec: %w", err)
	}

	return nil
}

func (r *repository) Get(ctx context.Context, in models.ExternalLovePercentage) (models.ExternalLovePercentage, error) {
	key := r.key(in)

	value, err := r.db.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return models.ExternalLovePercentage{}, ErrNotFound
	} else if err != nil {
		return models.ExternalLovePercentage{}, fmt.Errorf("get order: %w", err)
	}

	var lv models.ExternalLovePercentage
	err = json.Unmarshal([]byte(value), &lv)
	if err != nil {
		return models.ExternalLovePercentage{}, fmt.Errorf("failed to decode numbers json: %w", err)
	}

	return lv, nil
}

func (r *repository) Update(ctx context.Context, in models.ExternalLovePercentage) error {
	data, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to encode order: %w", err)
	}

	key := r.key(in)

	err = r.db.SetXX(ctx, key, string(data), 0).Err()
	if errors.Is(err, redis.Nil) {
		return ErrNotFound
	} else if err != nil {
		return fmt.Errorf("set order: %w", err)
	}

	return nil
}
