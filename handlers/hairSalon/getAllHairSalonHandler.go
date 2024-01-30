package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler pour la récupération de tous les salons de coiffure
func GetAllHairSalonHandler(c *gin.Context) {
	// Accédez à la connexion à la base de données
	db := services.GetConnection()

	// Déclarez une slice pour stocker tous les salons de coiffure
	var hairSalons []models.HairSalon

	// Effectuez une requête pour récupérer tous les salons de coiffure de la base de données
	if err := db.Find(&hairSalons).Error; err != nil {
		// En cas d'erreur lors de la recherche, renvoyez une réponse d'erreur
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hairSalons"})
		return
	}

	// Répondez avec les salons de coiffure récupérés
	c.JSON(http.StatusOK, gin.H{"hairSalons": hairSalons})
}
