package models

import (
	validators "GolandProject/validators/user"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Username  string `gorm:"unique;not null" json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// Number            int64     `json:"number"`
	Email             string `gorm:"unique;not null" json:"email"`
	EncryptedPassword string `gorm:"-" json:"-"`
	Password          string `json:"password"`
	// Roles             []string  `gorm:"-" json:"roles"`
	// BoutiqueID        int64     `json:"boutiqueID"`
	Model
}

func (u *User) HashPassword() ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
}

func (u *User) Create(CreateUser validators.RegisterUserValidator) {
	u.Username = CreateUser.Username
	u.FirstName = CreateUser.FirstName
	u.LastName = CreateUser.LastName
	u.Email = CreateUser.Email

	// If a new password is provided, update the password
	if CreateUser.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(CreateUser.Password), bcrypt.DefaultCost)
		u.EncryptedPassword = string(hashedPassword)
	}
}
