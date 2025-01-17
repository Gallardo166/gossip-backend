package models

type User struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}
