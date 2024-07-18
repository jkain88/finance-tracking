package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/models"
	"github.com/jkain88/finance-tracking/pkg/services"
	"github.com/jkain88/finance-tracking/pkg/utils"
	"github.com/markbates/goth/gothic"
)

func AuthRoutes(router *gin.RouterGroup, userService *services.UserService) {
	router.GET("/auth/:provider", func(c *gin.Context) {
		q := c.Request.URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request.URL.RawQuery = q.Encode()
		gothic.BeginAuthHandler(c.Writer, c.Request)
	})

	router.GET("/auth/:provider/callback", func(c *gin.Context) {
		q := c.Request.URL.Query()
		q.Add("provider", c.Param("provider"))
		c.Request.URL.RawQuery = q.Encode()
		user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		var userObject *models.User
		userExists, err := userService.IsUserExists(user.Email)
		if err != nil {
			fmt.Println("ERROR", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if !userExists {
			userObject, err = userService.CreateUser(user.Email, c.Param("provider"))
			if err != nil {
				fmt.Println("ERROR", err)
				c.JSON(200, gin.H{"error": err})
				return
			}
		} else {
			userObject, err = userService.GetUserByEmail(user.Email)
			if err != nil {
				fmt.Println("ERROR", err)
				c.JSON(200, gin.H{"error": err})
				return
			}
		}

		fmt.Println(userObject.Email, userObject.CreatedAt)
		token, err := utils.GenerateJWT(*userObject)
		if err != nil {
			fmt.Println("ERROR", err)
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"token": token})
	})
}
