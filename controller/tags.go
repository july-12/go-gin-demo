package controller

import (
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"

	"github.com/gin-gonic/gin"
)

func TagCreate(c *gin.Context) {
	var tag models.Tag
	tag.Name = c.PostForm("name")
	tag.UserID = c.GetUint("currentUserId")
	database.DB.Create(&tag)
	c.JSON(http.StatusOK, tag)
}

func TagUpdate(c *gin.Context) {
	var tag models.Tag
	err := database.DB.Find(&tag, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusOK, "")
	}
	tag.Name = c.PostForm("name")
	database.DB.Save(&tag)
	c.JSON(http.StatusOK, tag)
}

func TagIndex(c *gin.Context) {
	var tags []models.Tag
	database.DB.Find(&tags)
	c.JSON(http.StatusOK, tags)
}

func TagsIndexByPost(c *gin.Context) {
	var comments []models.Comment
	database.DB.Where("postId = ?", c.Param("id")).Find(&comments)
	c.JSON(http.StatusOK, comments)
}
