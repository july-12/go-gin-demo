package models

type Comment struct {
	BaseModelWithAuthor
	Content string `json:"title" gorm:"text;not null"`
	PostID  uint   `json:"-" gorm:"not null"`
}
