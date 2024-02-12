package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Structure de validation pour la mise à jour d'une réservation
type UpdateReservationValidator struct {
	DateOfReservation time.Time `json:"dateOfReservation" validate:"required"`
	Status            string    `json:"status" validate:"omitempty,min=2"`
	SalonID           int       `json:"salonId" validate:"required"`
	HairDresserID     int       `json:"HairDresserID" validate:"required"`
	UserID            int       `json:"UserID" validate:"required"`
}

// Méthode pour valider la structure UpdateReservationValidator
func (u *UpdateReservationValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
