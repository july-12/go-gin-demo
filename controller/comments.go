package controller

import (
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CommentFormInput struct {
	Content string `form:"content" binding:"required" json:"Content"`
}

func CommentCreate(c *gin.Context) {
	var input CommentFormInput
	err := c.BindJSON(&input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
	}
	var comment models.Comment
	comment.Content = input.Content
	postId, err := strconv.ParseUint(c.Param("id"), 0, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"miss postId": err})
		return
	}
	comment.PostID = uint(postId)
	comment.UserID = c.GetUint("currentUserId")
	database.DB.Create(&comment)
	c.JSON(http.StatusOK, comment)
}

func CommentShow(c *gin.Context) {
	var comments []models.Comment
	database.DB.Where("post_id = ?", c.Param("id")).Preload("Anchor").Find(&comments)
	c.JSON(http.StatusOK, comments)
}
