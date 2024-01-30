package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler pour la récupération d'un hairSalon par son ID
func GetSingleHairSalonHandler(c *gin.Context) {
	// Accédez à la connexion à la base de données
	db := services.GetConnection()

	// Déclarez une variable pour stocker le salon de coiffure
	var hairSalon models.HairSalon

	// Récupérez l'ID du paramètre de la requête
	hairSalonId := c.Param("hairSalonId")

	// Effectuez une requête pour récupérer le salon de coiffure de la base de données
	if err := db.First(&hairSalon, hairSalonId).Error; err != nil {
		// En cas d'erreur lors de la recherche, renvoyez une réponse d'erreur
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hairSalon"})
		return
	}

	// Répondez avec le salon de coiffure récupéré
	c.JSON(http.StatusOK, gin.H{"hairSalon": hairSalon})
}
