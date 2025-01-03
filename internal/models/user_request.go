package models

//@see https://github.com/go-playground/validator

type UserLogin struct {
	Account  string `json:"account" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=20"`
}
