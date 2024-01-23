package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"
	"GolandProject/validators"

	"github.com/gin-gonic/gin"
)

// Handler pour la mise à jour d'un utilisateur
func UpdateUserHandler(c *gin.Context) {
	// Récupérez l'ID de l'utilisateur à partir des paramètres de l'URL
	userID := c.Param("userId")

	// Accédez à la connexion à la base de données
	db := services.GetConnection()

	// Récupérez l'utilisateur à partir de la base de données en utilisant son ID
	var existingUser models.User
	if err := db.First(&existingUser, userID).Error; err != nil {
		// Affichez l'erreur et renvoyez une réponse NotFound
		fmt.Println("Error retrieving user from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// Affichez des informations pour débuguer
	fmt.Println("Existing User:", existingUser)

	// Accédez à l'utilisateur à partir du contexte
	user, exists := c.Get("user")
	if !exists {
		// Affichez un message en cas d'erreur
		fmt.Println("User information not available for this request put")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not available for this request put"})
		return
	}

	// Assurez-vous de convertir l'utilisateur en type *models.User
	updateUser, ok := user.(*models.User)
	if !ok {
		// Affichez un message en cas d'erreur
		fmt.Println("User information not available in the expected format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not available in the expected format"})
		return
	}

	var userValidator validators.UpdateUserValidator

	// Lisez et validez les données JSON envoyées par l'utilisateur
	if err := c.ShouldBindJSON(&userValidator); err != nil {
		// Affichez l'erreur de validation et renvoyez une réponse BadRequest
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validez la structure UpdateUserValidator
	if err := userValidator.Validate(); err != nil {
		// Affichez l'erreur de validation et renvoyez une réponse BadRequest
		fmt.Println("Error validating user input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Effectuez la mise à jour de l'utilisateur
	updateUser.Update(userValidator)

	// Enregistrez l'utilisateur mis à jour dans la base de données
	db.Save(updateUser)

	// Répondez avec les données de l'utilisateur mis à jour
	c.JSON(http.StatusOK, gin.H{"user": updateUser})
}
