package models

import "time"

type BaseModel struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type BaseModelWithAuthor struct {
	BaseModel
	UserID uint `json:"-"`
	Author User `json:"-" gorm:"foreignKey:UserID"` // `json:"-"`
}
