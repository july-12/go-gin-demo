package models

type Post struct {
	BaseModelWithAuthor
	Title    string    `json:"title" gorm:"size:100;not null"`
	Content  string    `json:"content" gorm:"text;not null"`
	Comments []Comment `json:"-"`

	Tags []Tag `json:"tags" gorm:"many2many:post_tags"`

	UserID uint `json:"-"`
	Anchor User `json:"anchor" gorm:"foreignKey:UserID"`
}
