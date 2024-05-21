package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/middlewares"
	"github.com/jkain88/finance-tracking/pkg/services"
)

func UserRoutes(router *gin.RouterGroup, service *services.UserService) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", service.CreateUser)
		userGroup.POST("/signin", service.SignIn)

		authenticated := userGroup.Group("/")
		authenticated.Use(middlewares.Authenticate)
		{
			authenticated.GET("me", service.Me)
			authenticated.GET("categories", service.UserCategories)
			authenticated.GET("accounts", service.UserAccounts)
		}
	}
	router.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Hello": "Test"})
	})
}
