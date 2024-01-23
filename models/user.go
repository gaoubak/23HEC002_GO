package models

import (
	"GolandProject/validators"

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
