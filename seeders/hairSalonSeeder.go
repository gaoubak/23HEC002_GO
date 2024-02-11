package seeders

import (
	"GolandProject/models"

	"gorm.io/gorm"
)

// HairSalonSeeder seeder pour les salons de coiffure
func HairSalonSeeder(db *gorm.DB) {
	// Créez des coiffeurs si aucun n'existe
	var existingCoiffeurs []models.HairDresser
	db.Find(&existingCoiffeurs)
	if len(existingCoiffeurs) == 0 {
		coiffeurs := []models.HairDresser{
			{
				Name:        "Mia (Senior Stylist)",
				Email:       "Mia@gmail.com",
				Speciality:  "Permanente",
				Description: "Senior styliste spécialisée dans la permanente et le lissage coréen.",
			},
			{
				Name:        "Elise (Senior Stylist)",
				Email:       "Elise@gmail.com",
				Speciality:  "Couleur",
				Description: "Senior styliste spécialisée dans la couleur et la décoloration.",
			},
			{
				Name:        "Taylor (Senior Stylist)",
				Email:       "Taylor@gmail.com",
				Speciality:  "Coupe dégradée",
				Description: "Senior styliste spécialisée dans la coupe dégradée et le lissage japonais.",
			},
		}
		db.Create(&coiffeurs)
		existingCoiffeurs = coiffeurs
	}

	// Créez des salons de coiffure
	salons := []models.HairSalon{
		{
			Name:        "Bleu coiffure",
			Adress:      "16 Rue d'Ouessant, 75015 Paris, France",
			Phone:       "0102030405",
			Email:       "SalonNum1@mail.com",
			Description: "Venez découvrir Bleu Coiffure, un salon de coiffure coréen de renom à Paris 15, spécialisé dans les techniques de pointe telles que la permanente digitale coréenne, le lissage japonais et lissage coréen, ainsi que la décoloration et la coloration capillaire.",
		},
		{
			Name:        "Hair Studio Greet",
			Adress:      "90 Rue de Richelieu, 75002 Paris, France",
			Phone:       "0102030406",
			Email:       "SalonNum2@mail.com",
			Description: "Salon de coiffure, coréen et japonais situé à Paris II, ​Hair Studio Greet ! Nous sommes votre destination incontournable pour des transformations capillaires",
		},
		{
			Name:        "Y Salon",
			Adress:      "6 Rue de la Coutellerie, 75004 Paris",
			Phone:       "0102030407",
			Email:       "SalonNum3@mail.com",
			Description: "Salon de coiffure Y Salon situé à Paris 4ème, spécialisé dans les techniques de coiffure française. Venez découvrir notre salon de coiffure et nos coiffeurs professionnels.",
		},
	}
	db.Create(&salons)

	// Associez les coiffeurs aux salons de coiffure
	for i, salon := range salons {
		db.Model(&salon).Association("HairDressers").Append(&existingCoiffeurs[i])
	}
}

