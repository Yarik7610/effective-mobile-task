package postgres

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/Yarik7610/effective-mobile-task/internal/dto"
	"github.com/Yarik7610/effective-mobile-task/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubscriptionRepository interface {
	CreateSubscription(ctx context.Context, subscription *model.Subscription) error
	ReadSubscription(ctx context.Context, subscriptionID string) (*model.Subscription, error)
	UpdateSubscription(ctx context.Context, updateSubscriptionDTO *dto.UpdateSubscription, updatedSubscription *model.Subscription, subscriptionID string) error
	DeleteSubscription(ctx context.Context, subscriptionID string) error
	ListSubscriptions(ctx context.Context, page, count uint, sort, order string) ([]model.Subscription, error)
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
	sql := `
		SELECT subscription_id, service_name, price, user_id, start_date, end_date
		FROM subscriptions
		WHERE subscription_id = $1
	`

	subscription := model.Subscription{}

	err := r.pool.QueryRow(ctx, sql, subscriptionID).Scan(
		&subscription.SubscriptionID,
		&subscription.ServiceName,
		&subscription.Price,
		&subscription.UserID,
		&subscription.StartDate,
		&subscription.EndDate,
	)

	if err != nil {
		return nil, err
	}

	return &subscription, nil
}

func (r *subscriptionRepository) UpdateSubscription(
	ctx context.Context,
	updateSubscriptionDTO *dto.UpdateSubscription,
	updatedSubscription *model.Subscription,
	subscriptionID string,
) error {
	setParts := []string{}
	args := []any{}
	idx := 1

	if updateSubscriptionDTO.ServiceName != nil {
		setParts = append(setParts, fmt.Sprintf("service_name = $%d", idx))
		args = append(args, *updateSubscriptionDTO.ServiceName)
		idx++
	}
	if updateSubscriptionDTO.Price != nil {
		setParts = append(setParts, fmt.Sprintf("price = $%d", idx))
		args = append(args, *updateSubscriptionDTO.Price)
		idx++
	}
	if updateSubscriptionDTO.StartDate != nil {
		setParts = append(setParts, fmt.Sprintf("start_date = $%d", idx))
		args = append(args, updatedSubscription.StartDate)
		idx++
	}
	if updateSubscriptionDTO.EndDate != nil {
		log.Println("HEELO")
		setParts = append(setParts, fmt.Sprintf("end_date = $%d", idx))
		args = append(args, updatedSubscription.EndDate)
		idx++
	}

	if len(setParts) == 0 {
		return fmt.Errorf("no rows to update were chosen")
	}

	args = append(args, subscriptionID)

	sql := fmt.Sprintf(`
		UPDATE subscriptions
		SET %s 
		WHERE subscription_id = $%d 
    RETURNING subscription_id, service_name, price, user_id, start_date, end_date
	`,
		strings.Join(setParts, ", "),
		idx,
	)

	err := r.pool.QueryRow(ctx, sql, args...).Scan(
		&updatedSubscription.SubscriptionID,
		&updatedSubscription.ServiceName,
		&updatedSubscription.Price,
		&updatedSubscription.UserID,
		&updatedSubscription.StartDate,
		&updatedSubscription.EndDate,
	)

	return err
}

func (r *subscriptionRepository) DeleteSubscription(ctx context.Context, subscriptionID string) error {
	sql := `
		DELETE FROM subscriptions
		WHERE subscription_id = $1
	`
	_, err := r.pool.Exec(ctx, sql, subscriptionID)
	return err
}

func (r *subscriptionRepository) ListSubscriptions(ctx context.Context, page, count uint, sort, order string) ([]model.Subscription, error) {
	sql := fmt.Sprintf(`
		SELECT subscription_id, service_name, price, user_id, start_date, end_date 
		FROM subscriptions
		ORDER BY %s %s 
		LIMIT $1
		OFFSET $2
	`, sort, order)

	rows, err := r.pool.Query(ctx, sql, count, (page-1)*count)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subscriptions []model.Subscription
	for rows.Next() {
		var subscription model.Subscription

		err := rows.Scan(
			&subscription.SubscriptionID,
			&subscription.ServiceName,
			&subscription.Price,
			&subscription.UserID,
			&subscription.StartDate,
			&subscription.EndDate,
		)
		if err != nil {
			return nil, err
		}

		subscriptions = append(subscriptions, subscription)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return subscriptions, nil
}
