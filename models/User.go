package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Password string `json:"-"`

	// Posts []Post `json:"-"`
}
