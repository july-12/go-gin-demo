package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Content string `json:"title" gorm:"text;not null"`
	PostID  uint   `json:"-" gorm:"not null"`
	UserID  uint   `json:"-" gorm:"not null"`
}
