package models

import (
	validators "GolandProject/validators/user"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                int    `gorm:"primaryKey" json:"id"`
	Username          string `gorm:"unique;not null" json:"username"`
	FirstName         string `json:"firstName"`
	LastName          string `json:"lastName"`
	Email             string `gorm:"unique;not null" json:"email"`
	EncryptedPassword string `gorm:"-" json:"-"`
	Password          string `json:"password"`
	// Model
}

// Méthode pour mettre à jour un utilisateur
func (u *User) Update(updateUser validators.UpdateUserValidator) {
	u.Username = updateUser.Username
	u.FirstName = updateUser.FirstName
	u.LastName = updateUser.LastName
	u.Email = updateUser.Email

	// Si un nouveau mot de passe est fourni, mettez à jour le mot de passe
	if updateUser.Password != "" {
		hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(updateUser.Password), bcrypt.DefaultCost)
		u.EncryptedPassword = string(hashedPassword)
	}
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
