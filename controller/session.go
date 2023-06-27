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
	Name     string `form:"name" json:"name"`
	Phone    string `form:"name" json:"phone"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Signup(c *gin.Context) {
	var input formInput
	var user models.User
	if err := c.BindJSON(&input); err == nil {
		user.Name = input.Name
		user.Phone = input.Phone
		password := []byte(input.Password)
		if passowrdHash, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost); err != nil {
			log.Println("err: ", err)
			c.JSON(http.StatusBadRequest, gin.H{"BcryptFail": err})
		} else {
			user.Password = string(passowrdHash)
			database.DB.Create(&user)
			if token, err := utils.GenerateToken(user.ID); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"JwtFail": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"token": token})
				return
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}
}

func Login(c *gin.Context) {
	var input formInput
	if err := c.BindJSON(&input); err == nil {
		var user models.User
		result := database.DB.Find(&user, "name = ?", input.Name)
		if result.Error != nil {
			c.JSON(http.StatusOK, gin.H{"user is unexist!": result.Error})
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
			c.JSON(http.StatusOK, gin.H{"wrong password!": err})
		} else {
			if token, err := utils.GenerateToken(user.ID); err != nil {
				c.String(http.StatusNoContent, "")
			} else {
				c.JSON(http.StatusOK, gin.H{"token": token})
			}
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}
}

func Logout(c *gin.Context) {

}
