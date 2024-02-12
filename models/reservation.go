package models

import (
	validators "GolandProject/validators/reservation"
	"time"
)

type Reservation struct {
	ID                int         `gorm:"primaryKey" json:"id"`
	DateOfReservation time.Time   `json:"dateOfReservation"`
	Status            string      `json:"status"`
	SalonID           int         `gorm:"foreignKey:SalonID" json:"SalonID"`
	HairDresserID     int         `gorm:"foreignKey:HairDresserId" json:"HairDresserId"`
	UserID            int         `gorm:"foreignKey:UserID" json:"UserId"`
	HairSalon         HairSalon   `gorm:"foreignKey:SalonID"`
	User              User        `gorm:"foreignKey:UserID"`
	HairDresser       HairDresser `gorm:"foreignKey:HairDresserID"`
}

func (r *Reservation) Create(createReservation validators.CreateReservationValidator) {
	r.DateOfReservation = createReservation.DateOfReservation
	r.Status = createReservation.Status
	r.SalonID = createReservation.SalonID
	r.HairDresserID = createReservation.HairDresserID
	r.UserID = createReservation.UserID
}

func (r *Reservation) Update(updateReservation validators.UpdateReservationValidator) {
	r.DateOfReservation = updateReservation.DateOfReservation
	r.Status = updateReservation.Status
	r.SalonID = updateReservation.SalonID
	r.HairDresserID = updateReservation.HairDresserID
	r.UserID = updateReservation.UserID
}
