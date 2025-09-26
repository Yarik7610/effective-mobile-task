package http

import (
	"log/slog"
	"net/http"

	"github.com/Yarik7610/effective-mobile-task/internal/dto"
	"github.com/Yarik7610/effective-mobile-task/internal/service"
	"github.com/gin-gonic/gin"
)

type SubscriptionController interface {
	CreateSubscription(ctx *gin.Context)
	ReadSubscription(ctx *gin.Context)
	UpdateSubscription(ctx *gin.Context)
	DeleteSubscription(ctx *gin.Context)
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

	subscription, err := c.subscriptionService.ReadSubscription(subscriptionID)
	if err != nil {

		slog.Error("Subscription read failed", slog.Any("error", err))
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, subscription)
}

func (c *subscriptionController) UpdateSubscription(ctx *gin.Context) {

}
func (c *subscriptionController) DeleteSubscription(ctx *gin.Context) {

}
