package models

type SignupUser struct {
	Username        string `json:"username" validate:"required,min=5,max=100"`
	Password        string `json:"password" validate:"required,min=8,max=100"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,eqfield=Password"`
}

type LoginUser struct {
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}

type User struct {
	Id       int    `json:"id"`
	Username string `json:"username" validate:"required,min=5"`
	Password string `json:"password" validate:"required,min=8"`
}
