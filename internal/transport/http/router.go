package http

import (
	"github.com/Yarik7610/effective-mobile-task/internal/repository/postgres"
	"github.com/Yarik7610/effective-mobile-task/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitRouter(pool *pgxpool.Pool) *gin.Engine {
	router := gin.Default()

	subscriptionRepository := postgres.NewSubsciptionRepository(pool)
	subscriptionService := service.NewSubsciptionService(subscriptionRepository)
	subscriptionController := NewSubsciptionController(subscriptionService)

	subscriptionsGroup := router.Group("/subscriptions")
	{
		subscriptionsGroup.GET("/:subscriptionID", subscriptionController.ReadSubscription)
		subscriptionsGroup.POST("/", subscriptionController.CreateSubscription)
		subscriptionsGroup.PUT("/:subscriptionID", subscriptionController.UpdateSubscription)
		subscriptionsGroup.DELETE("/:subscriptionID", subscriptionController.DeleteSubscription)
	}

	return router
}
