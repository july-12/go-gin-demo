package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"title" gorm:"text;not null"`
	PostId  uint   `json:"-"`
	// Post    Post
}
