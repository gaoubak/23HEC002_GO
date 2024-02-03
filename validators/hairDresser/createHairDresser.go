package validators

import (
	"github.com/go-playground/validator/v10"
)

// Structure de validation pour la création d'une réservation
type CreateHairDresserValidator struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Speciality  string `json:"speciality"`
	Description string `json:"description"`
	//ReservationID int    `gorm:"foreignKey:ReservationID" json:"reservationId"`
}

// Méthode pour valider la structure CreateHairDresserValidator
func (c *CreateHairDresserValidator) Validate() error {

	validate := validator.New()
	return validate.Struct(c)
}
