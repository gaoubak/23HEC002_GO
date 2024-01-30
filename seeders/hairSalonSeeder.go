package seeders

import (
	"GolandProject/models"

	"gorm.io/gorm"
)

// HairSalonSeeder seeder pour les salons de coiffure
func HairSalonSeeder(db *gorm.DB) {
	db.Create(&models.HairSalon{
		Name:        "Salon de coiffure 1",
		Adress:      "1 rue du salon",
		Phone:       "0102030405",
		Email:       "SalonNum1@mail.com",
		Description: "Salon de coiffure numéro 1",
	})
	db.Create(&models.HairSalon{
		Name:        "Salon de coiffure 2",
		Adress:      "2 rue du salon",
		Phone:       "0102030406",
		Email:       "SalonNum2@mail.com",
		Description: "Salon de coiffure numéro 2",
	})
	db.Create(&models.HairSalon{
		Name:        "Salon de coiffure 3",
		Adress:      "3 rue du salon",
		Phone:       "0102030407",
		Email:       "SalonNum3@mail.com",
		Description: "Salon de coiffure numéro 3",
	})
}
