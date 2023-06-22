package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	log.Println("=========", token)
	if len(token) == 0 {
		log.Println("user not login in")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized!"})
		c.Abort()
		return
	}
	c.Next()
}
