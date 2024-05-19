package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jkain88/finance-tracking/pkg/utils"
)

func Authenticate(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")

	if token == "" {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	claims, validated := utils.ValidateJWT(token)
	if !validated {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Not authorized"})
		return
	}

	c.Set("userId", claims.UserId)
	c.Next()
}
