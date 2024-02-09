package models

type HairSalon struct {
	ID           int           `gorm:"primaryKey" json:"id"`
	Name         string        `gorm:"unique;not null" json:"name"`
	Adress       string        `json:"adress"`
	Phone        string        `json:"phone"`
	Email        string        `json:"email"`
	Description  string        `json:"description"`
	HairDressers []HairDresser `gorm:"many2many:hairdresser_hairsalon;" json:"hairDressers"`
	Reservations []Reservation `gorm:"foreignKey:SalonID" json:"reservations"`
	// Model
}
