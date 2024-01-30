package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Structure de validation pour la création d'une réservation
type CreateReservationValidator struct {
	DateOfReservation time.Time `json:"dateOfReservation" validate:"required"`
	Status            string    `json:"status" validate:"omitempty,min=2"`
	SalonID           int       `json:"salonId" validate:"required"`
	HairdresserID     int       `json:"hairdresserId" validate:"required"`
	ClientID          int       `json:"clientId" validate:"required"`
}

// SetDefaultStatus sets the default value for the Status field
func (c *CreateReservationValidator) SetDefaultStatus() {
	if c.Status == "" {
		c.Status = "En Attente"
	}
}

// Méthode pour valider la structure CreateReservationValidator
func (c *CreateReservationValidator) Validate() error {
	// Set the default status before validating
	c.SetDefaultStatus()

	validate := validator.New()
	return validate.Struct(c)
}
