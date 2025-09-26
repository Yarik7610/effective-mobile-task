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

func (r *subscriptionRepository) CreateSubscription(ctx context.Context, subscription *model.Subscription) error {
	sql := `
		INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING subscription_id
	`

	var subscriptionID string

	err := r.pool.QueryRow(ctx, sql,
		subscription.ServiceName,
		subscription.Price,
		subscription.UserID,
		subscription.StartDate,
		subscription.EndDate,
	).Scan(&subscriptionID)

	if err != nil {
		return err
	}

	subscription.SubscriptionID = subscriptionID

	return nil
}
func (r *subscriptionRepository) ReadSubscription(ctx context.Context, subscriptionID string) (*model.Subscription, error) {
	return nil, nil
}
func (r *subscriptionRepository) PutSubscription(ctx context.Context, updatedSubscription *model.Subscription) error {
	return nil
}
func (r *subscriptionRepository) DeleteSubscription(ctx context.Context, subscriptionID string) error {
	return nil
}
