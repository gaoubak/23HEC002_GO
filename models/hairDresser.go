package models

import (
	validators "GolandProject/validators/hairDresser"
)

type HairDresser struct {
	ID           int           `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name"`
	Email        string        `json:"email"`
	Speciality   string        `json:"speciality"`
	Description  string        `json:"description"`
	Reservations []Reservation `gorm:"many2many:hairdresser_reservations;" json:"reservations"`
	/*Reservation   Reservation `gorm:"foreignKey:ReservationID"`*/
}

func (r *HairDresser) Create(createHairDresser validators.CreateHairDresserValidator) {
	r.Name = createHairDresser.Name
	r.Email = createHairDresser.Email
	r.Speciality = createHairDresser.Speciality
	r.Description = createHairDresser.Description
	//r.ReservationID = createHairDresser.ReservationID
}

func (r *HairDresser) Update(updateHairDresser validators.UpdateHairDresserValidator) {
	r.Name = updateHairDresser.Name
	r.Email = updateHairDresser.Email
	r.Speciality = updateHairDresser.Speciality
	r.Description = updateHairDresser.Description
	//r.ReservationID = updateHairDresser.ReservationID
}
