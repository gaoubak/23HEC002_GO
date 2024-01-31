package models

type HairDresser struct {
	ID            int    `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	Email 		  string `json:"email"`
	Speciality    string `json:"speciality"`
	Description   string `json:"description"`
	ReservationID []int  `gorm:"foreignKey:ReservationID" json:"reservationId"`
}
