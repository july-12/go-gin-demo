package controller

import (
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"
	"starter-with-docker/utils"

	"github.com/gin-gonic/gin"
)

func UserIndex(c *gin.Context) {
	var users []models.User
	database.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}

func UserShow(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	database.DB.Find(&user, id)
	// database.DB.Preload("Posts").Find(&user, id)
	// database.DB.Preload(clause.Associations).Find(&user, id)
	// database.DB.Model(&models.User{}).Preload("Posts").Find(&user, id)
	c.JSON(http.StatusOK, user)
}

func UserUpdate(c *gin.Context) {

}

func UserDelete(c *gin.Context) {

}

func CurrentUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.String(http.StatusOK, "")
		return
	}
	claims, err := utils.ValidateToken(token)
	if err != nil {
		c.String(http.StatusOK, "")
		return
	}
	var user models.User
	database.DB.Where("ID = ?", claims.UserId).First(&user)
	c.JSON(http.StatusOK, user)
}
