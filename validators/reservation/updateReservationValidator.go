package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Structure de validation pour la mise à jour d'une réservation
type UpdateReservationValidator struct {
	DateOfReservation time.Time `json:"dateOfReservation" validate:"required"`
	SalonID           int       `json:"salonId" validate:"required"`
	HairdresserID     int       `json:"hairdresserId" validate:"required"`
	ClientID          int       `json:"clientId" validate:"required"`
}

// Méthode pour valider la structure UpdateReservationValidator
func (u *UpdateReservationValidator) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
