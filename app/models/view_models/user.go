package view_models

type User struct {
	Base
	UserName string `json:"username"`
	Email    string `json:"email"`
}
