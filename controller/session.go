package controller

import (
	"log"
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"
	"starter-with-docker/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type formInput struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input formInput
	var user models.User
	if err := c.BindJSON(&input); err == nil {
		user.Name = input.Name
		password := []byte(input.Password)
		if passowrdHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost); err != nil {
			log.Println("err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"BcryptFail": err})
		} else {
			user.Password = string(passowrdHash)
			database.DB.Create(&user)
			if token, err := utils.GenerateToken(user.ID, passowrdHash); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"JwtFail": err})
			} else {
				c.JSON(http.StatusOK, token)
				return

			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}
}

func Login(c *gin.Context) {
	var input formInput
	c.BindJSON(&input)
}

func Logout(c *gin.Context) {

}
