package models

type Comment struct {
	BaseModelWithAuthor
	Content string `json:"content" gorm:"text;not null"`
	PostID  uint   `json:"-" gorm:"not null"`
	UserID  uint   `json:"-" gorm:"not null"`
	Anchor  User   `json:"anchor" gorm:"foreignKey:UserID"`
}
