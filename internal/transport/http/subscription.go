package http

import (
	"log/slog"
	"net/http"

	"github.com/Yarik7610/effective-mobile-task/internal/dto"
	"github.com/Yarik7610/effective-mobile-task/internal/query"
	"github.com/Yarik7610/effective-mobile-task/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SubscriptionController interface {
	CreateSubscription(ctx *gin.Context)
	ReadSubscription(ctx *gin.Context)
	UpdateSubscription(ctx *gin.Context)
	DeleteSubscription(ctx *gin.Context)
	ListSubscriptions(ctx *gin.Context)
	TotalSubscriptionsPrice(ctx *gin.Context)
}

type subscriptionController struct {
	subscriptionService service.SubscriptionService
}

func NewSubsciptionController(subscriptionService service.SubscriptionService) SubscriptionController {
	return &subscriptionController{subscriptionService: subscriptionService}
}

func (c *subscriptionController) CreateSubscription(ctx *gin.Context) {
	var createSubscriptionDTO dto.CreateSubscription

	if err := ctx.ShouldBindJSON(&createSubscriptionDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscription, customErr := c.subscriptionService.CreateSubscription(&createSubscriptionDTO)
	if customErr != nil {
		slog.Error("Subscription create failed", slog.Any("error", customErr))
		ctx.JSON(customErr.Code, gin.H{"error": customErr.Message})
		return
	}

	ctx.JSON(http.StatusCreated, subscription)
}

func (c *subscriptionController) ReadSubscription(ctx *gin.Context) {
	subscriptionID := ctx.Param("subscriptionID")

	if _, err := uuid.Parse(subscriptionID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "subscriptionID isn't a valid UUID"})
		return
	}

	subscription, err := c.subscriptionService.ReadSubscription(subscriptionID)
	if err != nil {
		slog.Error("Subscription read failed", slog.Any("error", err))
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, subscription)
}

func (c *subscriptionController) UpdateSubscription(ctx *gin.Context) {
	subscriptionID := ctx.Param("subscriptionID")

	if _, err := uuid.Parse(subscriptionID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "subscriptionID isn't a valid UUID"})
		return
	}

	var updateSubscriptionDTO dto.UpdateSubscription
	if err := ctx.ShouldBindJSON(&updateSubscriptionDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedSubscription, err := c.subscriptionService.UpdateSubscription(&updateSubscriptionDTO, subscriptionID)
	if err != nil {
		slog.Error("Subscription update failed", slog.Any("error", err))
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, updatedSubscription)
}

func (c *subscriptionController) DeleteSubscription(ctx *gin.Context) {
	subscriptionID := ctx.Param("subscriptionID")

	if _, err := uuid.Parse(subscriptionID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "subscriptionID isn't a valid UUID"})
		return
	}

	err := c.subscriptionService.DeleteSubscription(subscriptionID)
	if err != nil {
		slog.Error("Subscription delete failed", slog.Any("error", err))
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.AbortWithStatus(http.StatusNoContent)
}

func (c *subscriptionController) ListSubscriptions(ctx *gin.Context) {
	const DEFAULT_PAGINATE_COUNT = 5
	const DEFAULT_SORT_FIELD = "service_name"
	const DEFAULT_ORDER_FIELD = "asc"

	var listSubscriptionsQuery query.ListSubscriptions

	if err := ctx.ShouldBindQuery(&listSubscriptionsQuery); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if listSubscriptionsQuery.Count == 0 {
		listSubscriptionsQuery.Count = DEFAULT_PAGINATE_COUNT
	}
	if listSubscriptionsQuery.Sort == "" {
		listSubscriptionsQuery.Sort = DEFAULT_SORT_FIELD
	}
	if listSubscriptionsQuery.Order == "" {
		listSubscriptionsQuery.Order = DEFAULT_ORDER_FIELD
	}

	subscriptions, err := c.subscriptionService.ListSubscriptions(&listSubscriptionsQuery)
	if err != nil {
		slog.Error("Subscription listing failed", slog.Any("error", err))
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, subscriptions)
}

func (c *subscriptionController) TotalSubscriptionsPrice(ctx *gin.Context) {

}
