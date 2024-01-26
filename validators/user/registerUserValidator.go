package validators

import (
	"github.com/go-playground/validator/v10"
)

type RegisterUserValidator struct {
	Username  string `json:"username" validate:"required,min=2"`
	FirstName string `json:"firstName" validate:"required,min=2"`
	LastName  string `json:"lastName" validate:"required,min=2"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
}

// Validate method to validate the RegisterUserValidator structure
func (u *RegisterUserValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
