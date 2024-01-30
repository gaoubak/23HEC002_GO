package models

import (
	validators "GolandProject/validators/reservation"
	"time"
)

type Reservation struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	DateOfReservation time.Time `json:"dateOfReservation"`
	Status            string    `json:"status"`
	SalonID           int       `gorm:"foreignKey:SalonId"`
	HairdresserID     int       `gorm:"foreignKey:HairdresserId" json:"hairdresserId"`
	ClientID          int       `gorm:"foreignKey:ClientID" json:"clientId"`
	/*Salon             Salon     `gorm:"foreignKey:SalonId"`
	Hairdresser       User      `gorm:"foreignKey:HairdresserId"`*/
	Client User `gorm:"foreignKey:ClientID"`
}

func (r *Reservation) Create(createReservation validators.CreateReservationValidator) {
	r.DateOfReservation = createReservation.DateOfReservation
	r.Status = createReservation.Status
	r.SalonID = createReservation.SalonID
	r.HairdresserID = createReservation.HairdresserID
	r.ClientID = createReservation.ClientID
}

func (r *Reservation) Update(updateReservation validators.UpdateReservationValidator) {
	r.DateOfReservation = updateReservation.DateOfReservation
	r.Status = updateReservation.Status
	r.SalonID = updateReservation.SalonID
	r.HairdresserID = updateReservation.HairdresserID
	r.ClientID = updateReservation.ClientID
}
