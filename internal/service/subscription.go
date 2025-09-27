package service

import (
	"context"
	"database/sql"
	"errors"
	"net/http"
	"time"

	"github.com/Yarik7610/effective-mobile-task/internal/dto"
	"github.com/Yarik7610/effective-mobile-task/internal/model"
	"github.com/Yarik7610/effective-mobile-task/internal/query"
	"github.com/Yarik7610/effective-mobile-task/internal/repository/postgres"
	"github.com/Yarik7610/effective-mobile-task/utils"
)

var ErrSubscriptionNotFound = errors.New("subscription not found")

type SubscriptionService interface {
	CreateSubscription(createSubscriptionDTO *dto.CreateSubscription) (*model.Subscription, *utils.Err)
	ReadSubscription(subscriptionID string) (*model.Subscription, *utils.Err)
	UpdateSubscription(updateSubscriptionDTO *dto.UpdateSubscription, subscriptionID string) (*model.Subscription, *utils.Err)
	DeleteSubscription(subscriptionID string) *utils.Err
	ListSubscriptions(listSubscriptionsQuery *query.ListSubscriptions) ([]model.Subscription, *utils.Err)
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

func (s *subscriptionService) UpdateSubscription(updateSubscriptionDTO *dto.UpdateSubscription, subscriptionID string) (*model.Subscription, *utils.Err) {
	var startDateTime time.Time
	if updateSubscriptionDTO.StartDate != nil {
		t, err := utils.ParseMonthYearStringToTime(*updateSubscriptionDTO.StartDate)
		if err != nil {
			return nil, utils.NewErr(http.StatusBadRequest, err.Error())
		}
		startDateTime = t
	}

	var endDateTime *time.Time
	if updateSubscriptionDTO.EndDate != nil {
		t, err := utils.ParseMonthYearStringToTime(*updateSubscriptionDTO.EndDate)
		if err != nil {
			return nil, utils.NewErr(http.StatusBadRequest, err.Error())
		}
		endDateTime = &t
	}

	updatedSubscription := model.Subscription{
		SubscriptionID: subscriptionID,
		StartDate:      startDateTime,
		EndDate:        endDateTime,
	}

	err := s.subscriptionRepository.UpdateSubscription(context.Background(), updateSubscriptionDTO, &updatedSubscription, subscriptionID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, utils.NewErr(http.StatusNotFound, ErrSubscriptionNotFound.Error())
		}
		return nil, utils.NewErr(http.StatusInternalServerError, err.Error())
	}

	return &updatedSubscription, nil
}

func (s *subscriptionService) DeleteSubscription(subscriptionID string) *utils.Err {
	err := s.subscriptionRepository.DeleteSubscription(context.Background(), subscriptionID)
	if err != nil {
		return utils.NewErr(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func (s *subscriptionService) ListSubscriptions(listSubscriptionsQuery *query.ListSubscriptions) ([]model.Subscription, *utils.Err) {
	return nil, nil
}
