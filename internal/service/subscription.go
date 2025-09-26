package service

import (
	"github.com/Yarik7610/effective-mobile-task/internal/dto"
	"github.com/Yarik7610/effective-mobile-task/internal/model"
	"github.com/Yarik7610/effective-mobile-task/internal/repository/postgres"
	"github.com/Yarik7610/effective-mobile-task/utils"
)

type SubscriptionService interface {
	CreateSubscription(createSubscriptionDTO *dto.CreateSubscription) (*model.Subscription, *utils.Err)
	ReadSubscription(subscriptionID string) (*model.Subscription, *utils.Err)
	PutSubscription(createSubscriptionDTO *dto.UpdateSubscription) (*model.Subscription, *utils.Err)
	DeleteSubscription(subscriptionID string) *utils.Err
}

type subscriptionService struct {
	subscriptionRepository postgres.SubscriptionRepository
}

func NewSubsciptionService(subscriptionRepository postgres.SubscriptionRepository) SubscriptionService {
	return &subscriptionService{subscriptionRepository: subscriptionRepository}
}

func (c *subscriptionService) CreateSubscription(createSubscriptionDTO *dto.CreateSubscription) (*model.Subscription, *utils.Err) {
	return nil, nil
}
func (c *subscriptionService) ReadSubscription(subscriptionID string) (*model.Subscription, *utils.Err) {
	return nil, nil
}
func (c *subscriptionService) PutSubscription(createSubscriptionDTO *dto.UpdateSubscription) (*model.Subscription, *utils.Err) {
	return nil, nil
}
func (c *subscriptionService) DeleteSubscription(subscriptionID string) *utils.Err {
	return nil
}
