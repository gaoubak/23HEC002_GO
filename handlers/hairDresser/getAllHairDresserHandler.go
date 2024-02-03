package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler pour la récupération de tous les DresserHairDresser de coiffure
func GetAllHairDresserHandler(c *gin.Context) {
	// Accédez à la connexion à la base de données
	db := services.GetConnection()

	// Déclarez une slice pour stocker tous les HairDresser
	var HairDresser []models.HairDresser

	// Effectuez une requête pour récupérer tous les DresserHairDresser de coiffure de la base de données
	if err := db.Find(&HairDresser).Error; err != nil {
		// En cas d'erreur lors de la recherche, renvoyez une réponse d'erreur
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve HairDresser"})
		return
	}

	// Répondez avec les HairDresser de coiffure récupérés
	c.JSON(http.StatusOK, gin.H{"HairDresser": HairDresser})
}
