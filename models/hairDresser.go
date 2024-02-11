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
	HairSalons   []*HairSalon   `gorm:"many2many:hairdresser_hairsalon;" json:"hairSalons"`
	Reservations []Reservation `gorm:"many2many:hairdresser_reservations;" json:"reservations"`
}

func (r *HairDresser) Create(createHairDresser validators.CreateHairDresserValidator) {
	r.Name = createHairDresser.Name
	r.Email = createHairDresser.Email
	r.Speciality = createHairDresser.Speciality
	r.Description = createHairDresser.Description
}

func (r *HairDresser) Update(updateHairDresser validators.UpdateHairDresserValidator) {
	r.Name = updateHairDresser.Name
	r.Email = updateHairDresser.Email
	r.Speciality = updateHairDresser.Speciality
	r.Description = updateHairDresser.Description
}
