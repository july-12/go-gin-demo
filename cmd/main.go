package main

import (
	"net/http"
	database "starter-with-docker/config"
	"starter-with-docker/models"

	"github.com/gin-gonic/gin"
)

func main() {
	database.ConnectDb()
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello gin222"})
	})

	r.GET("/facts", func(c *gin.Context) {
		var facts []models.Fact
		database.DB.Find(&facts)
		c.JSON(http.StatusOK, facts)
	})

	r.POST("/fact", func(c *gin.Context) {
		var fact models.Fact
		c.BindJSON(&fact)
		database.DB.Create(&fact)
		var facts []models.Fact
		database.DB.Find(&facts)
		c.JSON(http.StatusOK, facts)
	})

	r.Run(":3000")
}
