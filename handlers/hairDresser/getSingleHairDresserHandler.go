package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler pour la récupération d'un hairDresser par son ID
func GetSingleHairDresserHandler(c *gin.Context) {
	// Accédez à la connexion à la base de données
	db := services.GetConnection()

	// Déclarez une variable pour stocker le Dresser de coiffure
	var hairDresser models.HairDresser

	// Récupérez l'ID du paramètre de la requête
	hairDresserId := c.Param("hairDresserId")

	// Effectuez une requête pour récupérer le Dresser de coiffure de la base de données
	if err := db.First(&hairDresser, hairDresserId).Error; err != nil {
		// En cas d'erreur lors de la recherche, renvoyez une réponse d'erreur
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve hairDresser"})
		return
	}

	// Répondez avec le Dresser de coiffure récupéré
	c.JSON(http.StatusOK, gin.H{"hairDresser": hairDresser})
}
