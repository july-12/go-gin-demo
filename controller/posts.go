package controller

import (
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"

	"github.com/gin-gonic/gin"
)

type postFormInput struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"Content"`
}

func PostCreate(c *gin.Context) {
	var input postFormInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}

	var post models.Post
	post.Title = input.Title
	post.Content = input.Content
	post.UserID = c.GetUint("currentUserId")
	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	database.DB.Find(&posts)
	// database.DB.Preload(clause.Associations).Find(&posts)

	c.JSON(http.StatusOK, posts)
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	database.DB.Find(&post, id)
	// database.DB.Preload("Author").Find(&post, id)
	c.JSON(http.StatusOK, post)
}
