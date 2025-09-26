package http

import (
	"github.com/Yarik7610/effective-mobile-task/internal/service"
	"github.com/gin-gonic/gin"
)

type SubscriptionController interface {
	CreateSubscription(ctx *gin.Context)
	ReadSubscription(ctx *gin.Context)
	PutSubscription(ctx *gin.Context)
	DeleteSubscription(ctx *gin.Context)
}

type subscriptionController struct {
	subscriptionService service.SubscriptionService
}

func NewSubsciptionController(subscriptionService service.SubscriptionService) SubscriptionController {
	return &subscriptionController{subscriptionService: subscriptionService}
}

func (c *subscriptionController) CreateSubscription(ctx *gin.Context) {

}
func (c *subscriptionController) ReadSubscription(ctx *gin.Context) {

}
func (c *subscriptionController) PutSubscription(ctx *gin.Context) {

}
func (c *subscriptionController) DeleteSubscription(ctx *gin.Context) {

}
