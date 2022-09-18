package models

type User struct {
	BaseModel
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string
	Role     string
}
