package postgres

import (
	"context"

	"github.com/Yarik7610/effective-mobile-task/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, subscription *model.Subscription) error
	ReadSubscription(ctx context.Context, subscriptionID string) (*model.Subscription, error)
	PutSubscription(ctx context.Context, updatedSubscription *model.Subscription) error
	DeleteSubscription(ctx context.Context, subscriptionID string) error
}

type subscriptionRepository struct {
	pool *pgxpool.Pool
}

func NewSubsciptionRepository(pool *pgxpool.Pool) SubscriptionRepository {
	return &subscriptionRepository{pool: pool}
}

func (c *subscriptionRepository) CreateSubscription(ctx context.Context, subscription *model.Subscription) error {
	return nil
}
func (c *subscriptionRepository) ReadSubscription(ctx context.Context, subscriptionID string) (*model.Subscription, error) {
	return nil, nil
}
func (c *subscriptionRepository) PutSubscription(ctx context.Context, updatedSubscription *model.Subscription) error {
	return nil
}
func (c *subscriptionRepository) DeleteSubscription(ctx context.Context, subscriptionID string) error {
	return nil
}
