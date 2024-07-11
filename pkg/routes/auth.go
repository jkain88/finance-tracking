package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/markbates/goth/gothic"
)

func AuthRoutes(router *gin.RouterGroup) {
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
			c.JSON(500, gin.H{"error": err})
			return
		}
		c.JSON(200, gin.H{"user": user})
	})
}
