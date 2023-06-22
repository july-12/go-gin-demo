package models

import (
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Title    string    `json:"title" gorm:"size:100;not null"`
	Content  string    `json:"content" gorm:"text;not null"`
	Comments []Comment `json:"-"`
	UserID   uint      `json:"-"`
	Author   User      `json:"-" gorm:"foreignKey:UserID"` // `json:"-"`
}
