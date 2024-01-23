package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserHandler(c *gin.Context) {
	// Accédez à l'utilisateur à partir du contexte
	user, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User information not available"})
		return
	}

	// Utilisez l'objet user comme nécessaire...
	c.JSON(http.StatusOK, gin.H{"user": user})
}
