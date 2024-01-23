package validators

import (
	"github.com/go-playground/validator/v10"
)

// Structure de validation pour la mise à jour d'un utilisateur
type UpdateUserValidator struct {
	Username  string `json:"username" validate:"omitempty,min=2"`
	FirstName string `json:"firstName" validate:"omitempty,min=2"`
	LastName  string `json:"lastName" validate:"omitempty,min=2"`
	Email     string `json:"email" validate:"omitempty,email"`
	Password  string `json:"password" validate:"omitempty,min=6"`
}

// Méthode pour valider la structure UpdateUserValidator
func (u *UpdateUserValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
