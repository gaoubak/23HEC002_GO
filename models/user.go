package models

import (
	"time"
)

type User struct {
	ID                int       `gorm:"primaryKey" json:"id"`
	Username          string    `gorm:"unique;not null" json:"username"`
	FirstName         string    `json:"firstName"`
	LastName          string    `json:"lastName"`
	Number            int64     `json:"number"`
	Email             string    `gorm:"unique;not null" json:"email"`
	EncryptedPassword string    `gorm:"-" json:"-"`
	Password          string    `json:"password"`
	Roles             []string  `gorm:"-" json:"roles"`
	CreatedAt         time.Time `json:"createdAt"`
	BoutiqueID        int64     `json:"boutiqueID"`
}
