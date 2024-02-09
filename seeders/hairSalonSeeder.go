package seeders

import (
	"GolandProject/models"

	"gorm.io/gorm"
)

// HairSalonSeeder seeder pour les salons de coiffure
func HairSalonSeeder(db *gorm.DB) {
	// Vérifiez si les coiffeurs existent déjà dans la base de données
	var existingCoiffeurs []models.HairDresser
	db.Find(&existingCoiffeurs)

	// Si aucun coiffeur n'existe, créez-en de nouveaux
	if len(existingCoiffeurs) == 0 {
		// Créez des coiffeurs...
		coiffeur1 := models.HairDresser{
			Name:        "Coiffeur 1",
			Email:       "Coiffeur1@gmail.com",
			Speciality:  "Permanente",
			Description: "Coiffeur numéro 1",
		}
		coiffeur2 := models.HairDresser{
			Name:        "Coiffeur 2",
			Email:       "Coiffeur2@gmail.com",
			Speciality:  "Couleur",
			Description: "Coiffeur numéro 2",
		}
		coiffeur3 := models.HairDresser{
			Name:        "Coiffeur 3",
			Email:       "Coiffeur3@gmail.com",
			Speciality:  "Coupe dégradée",
			Description: "Coiffeur numéro 3",
		}

		// Créez les enregistrements dans la base de données
		db.Create(&coiffeur1)
		db.Create(&coiffeur2)
		db.Create(&coiffeur3)
	}

	// Vérifiez si les salons de coiffure existent déjà dans la base de données
	var existingSalons []models.HairSalon
	db.Find(&existingSalons)

	// Si aucun salon de coiffure n'existe, créez-en de nouveaux
	if len(existingSalons) == 0 {
		// Créez des salons de coiffure...
		salon1 := models.HairSalon{
			Name:        "Salon de coiffure 1",
			Adress:      "1 rue du salon",
			Phone:       "0102030405",
			Email:       "SalonNum1@mail.com",
			Description: "Salon de coiffure numéro 1",
		}
		salon2 := models.HairSalon{
			Name:        "Salon de coiffure 2",
			Adress:      "2 rue du salon",
			Phone:       "0102030406",
			Email:       "SalonNum2@mail.com",
			Description: "Salon de coiffure numéro 2",
		}
		salon3 := models.HairSalon{
			Name:        "Salon de coiffure 3",
			Adress:      "3 rue du salon",
			Phone:       "0102030407",
			Email:       "SalonNum3@mail.com",
			Description: "Salon de coiffure numéro 3",
		}

		// Créez les enregistrements dans la base de données
		db.Create(&salon1)
		db.Create(&salon2)
		db.Create(&salon3)
	}
}
