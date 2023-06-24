package models

type User struct {
	BaseModel
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Password string `json:"-"`
}
