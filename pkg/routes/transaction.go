package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/middlewares"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func TransactionRoutes(route *gin.RouterGroup, service *services.TransactionService) {
	transactionGroup := route.Group("/transaction")
	{
		authenticated := transactionGroup.Group("/")
		authenticated.Use(middlewares.Authenticate)

		authenticated.POST("/create")
	}
}
