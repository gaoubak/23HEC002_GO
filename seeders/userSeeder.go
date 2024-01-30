package seeders

import (
	"GolandProject/models"

	"gorm.io/gorm"
)

// UserSeeder seeder pour les utilisateurs
func UserSeeder(db *gorm.DB) {
	db.Create(&models.User{
		Username:  "Admin",
		FirstName: "Jean",
		LastName:  "Dupont",
		Password:  "admin",
		Email:     "admin@gmail.com",
	})
	db.Create(&models.User{
		Username:  "User",
		FirstName: "Pierre",
		LastName:  "Douteau",
		Password:  "password",
		Email:     "user@mail.com",
	})
	db.Create(&models.User{
		Username:  "User2",
		FirstName: "Paul",
		LastName:  "Douteau",
		Password:  "password",
		Email:     "user2@mail.com",
	})
}
