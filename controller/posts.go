package controller

import (
	"net/http"
	database "starter-with-docker/db"
	"starter-with-docker/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
)

type postFormInput struct {
	Title   string `form:"title" json:"title" binding:"required"`
	Content string `form:"content" json:"Content"`
	TagID   uint   `form:"tagId" json:"tagId"`
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

	// var tag models.Tag
	// database.DB.Find(&tag, input.TagID)
	// database.DB.Create(&post).Association("Tags").Append(&tag)
	database.DB.Create(&post)
	c.JSON(http.StatusOK, post)
}

type PostQuery struct {
	Page    int `form:"page,default=0"`
	PageNum int `form:"pageNum,default=10"`
}

func PostIndex(c *gin.Context) {
	var query PostQuery
	c.ShouldBindQuery(&query)
	var posts []models.Post
	// database.DB.Find(&posts)
	result := database.DB.Limit(query.PageNum).Offset(query.Page).Preload(clause.Associations).Find(&posts)

	c.JSON(http.StatusOK, gin.H{"list": posts, "total": result.RowsAffected})
}

func PostIndexOfUser(c *gin.Context) {

	var posts []models.Post
	userId := c.Param("id")
	database.DB.Where("user_id = ?", userId).Find(&posts)
	// database.DB.Preload(clause.Associations).Find(&posts)

	c.JSON(http.StatusOK, posts)
}

func PostShow(c *gin.Context) {
	id := c.Param("id")
	var post models.Post
	database.DB.Find(&post, id)
	database.DB.Preload("Tags").Find(&post, id)
	c.JSON(http.StatusOK, post)
}
