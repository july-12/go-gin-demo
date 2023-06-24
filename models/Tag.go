package models

type Tag struct {
	BaseModelWithAuthor
	Name string `json:"name"`
}
