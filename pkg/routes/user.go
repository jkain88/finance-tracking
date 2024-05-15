package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func UserRoutes(router *gin.RouterGroup, service *services.UserService) {
	userGroup := router.Group("/user")
	{
		userGroup.GET("/signup", service.CreateUser)
	}
	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "Test"})
	})
}
