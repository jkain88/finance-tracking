package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/middlewares"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func CategoryRoutes(router *gin.RouterGroup, service *services.CategoryService) {
	categoryGroup := router.Group("/category")
	{
		authenticated := categoryGroup.Group("/")
		authenticated.Use(middlewares.Authenticate)

		authenticated.POST("/create", service.CreateCategory)
		authenticated.PUT("/:id", service.UpdateCategory)
	}
}
