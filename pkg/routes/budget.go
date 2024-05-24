package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/middlewares"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func BudgetRoutes(router *gin.RouterGroup, service *services.BudgetService) {
	budgetGroup := router.Group("/budget")
	{
		authenticated := budgetGroup.Group("/")
		authenticated.Use(middlewares.Authenticate)

		authenticated.POST("/create", service.CreateBudget)
		authenticated.PUT("/:id", service.UpdateBudget)
		// authenticated.DELETE("/:id", service.DeleteBudget)
	}
}
