package controller

import (
	"net/http"
	database "starter-with-docker/config"
	"starter-with-docker/models"

	"github.com/gin-gonic/gin"
)

type signupInput struct {
	Name     string `form:"name" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
type loginInput struct {
	Name     string `form:"name" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input signupInput
	var user models.User
	if c.BindJSON(&input) == nil {
		user.Name = input.Name
		user.Password = input.Password
		database.DB.Create(&user)
		c.JSON(http.StatusOK, user)
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{"msg": "params wrong!"})
}

func Login(c *gin.Context) {
	var input loginInput
	c.BindJSON(&input)
}

func Logout(c *gin.Context) {

}
