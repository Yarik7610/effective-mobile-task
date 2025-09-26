package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Yarik7610/effective-mobile-task/internal/dto"
	"github.com/Yarik7610/effective-mobile-task/internal/model"
	"github.com/Yarik7610/effective-mobile-task/internal/repository/postgres"
	"github.com/Yarik7610/effective-mobile-task/utils"
)

var ErrSubscriptionNotFound = errors.New("subscription not found")

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

func (s *subscriptionService) CreateSubscription(createSubscriptionDTO *dto.CreateSubscription) (*model.Subscription, *utils.Err) {
	startDateTime, err := utils.ParseMonthYearStringToTime(createSubscriptionDTO.StartDate)
	if err != nil {
		return nil, utils.NewErr(http.StatusBadRequest, err.Error())
	}

	var endDateTime *time.Time
	if createSubscriptionDTO.EndDate != nil {
		t, err := utils.ParseMonthYearStringToTime(*createSubscriptionDTO.EndDate)
		if err != nil {
			return nil, utils.NewErr(http.StatusBadRequest, err.Error())
		}
		endDateTime = &t
	}

	subscription := model.Subscription{
		ServiceName: createSubscriptionDTO.ServiceName,
		Price:       createSubscriptionDTO.Price,
		UserID:      createSubscriptionDTO.UserID,
		StartDate:   startDateTime,
		EndDate:     endDateTime,
	}

	err = s.subscriptionRepository.CreateSubscription(context.Background(), &subscription)
	if err != nil {
		return nil, utils.NewErr(http.StatusInternalServerError, err.Error())
	}

	return &subscription, nil
}

func (s *subscriptionService) ReadSubscription(subscriptionID string) (*model.Subscription, *utils.Err) {
	subscription, err := s.subscriptionRepository.ReadSubscription(context.Background(), subscriptionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.NewErr(http.StatusNotFound, ErrSubscriptionNotFound.Error())
		}
		return nil, utils.NewErr(http.StatusInternalServerError, err.Error())
	}

	return subscription, nil
}

func (s *subscriptionService) PutSubscription(createSubscriptionDTO *dto.UpdateSubscription) (*model.Subscription, *utils.Err) {
	return nil, nil
}
func (s *subscriptionService) DeleteSubscription(subscriptionID string) *utils.Err {
	return nil
}
