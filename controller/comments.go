package controller

import (
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CommentCreate(c *gin.Context) {
	var comment models.Comment
	content := c.PostForm("content")
	comment.Content = content
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
	database.DB.Where("postId = ?", c.Param("id")).Find(&comments)
	c.JSON(http.StatusOK, comments)
}
