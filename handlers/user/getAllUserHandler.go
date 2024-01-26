package handlers

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "GolandProject/services"
    "GolandProject/models"
)

// Handler pour la récupération de tous les utilisateurs
func GetAllUserHandler(c *gin.Context) {
    // Accédez à la connexion à la base de données
    db := services.GetConnection()

    // Déclarez une slice pour stocker tous les utilisateurs
    var users []models.User

    // Effectuez une requête pour récupérer tous les utilisateurs de la base de données
    if err := db.Find(&users).Error; err != nil {
        // En cas d'erreur lors de la recherche, renvoyez une réponse d'erreur
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }

    // Répondez avec les utilisateurs récupérés
    c.JSON(http.StatusOK, gin.H{"users": users})
}