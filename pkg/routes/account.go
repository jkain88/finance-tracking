package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/middlewares"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func AccountRoutes(router *gin.RouterGroup, service *services.AccountService) {
	accountGroup := router.Group("/account")
	{
		authenticated := accountGroup.Group("/")
		authenticated.Use(middlewares.Authenticate)

		authenticated.POST("/create", service.CreateAccount)
		// authenticated.PUT("/:id", service.UpdateAccount)
		// authenticated.DELETE("/:id", service.DeleteAccount)
	}
}
