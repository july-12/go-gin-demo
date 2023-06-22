package middleware

import (
	"net/http"
	"starter-with-docker/utils"

	"github.com/gin-gonic/gin"
)

func AuthorizedRoute(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "token empty!"})
		c.Abort()
		return
	}
	claims, err := utils.ValidateToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized!"})
		c.Abort()
		return
	}
	c.Set("currentUserId", claims.UserId)
	c.Next()
}
