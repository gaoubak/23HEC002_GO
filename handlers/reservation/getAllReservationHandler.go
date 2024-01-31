package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllReservationsHandler retrieves all reservations from the database
func GetAllReservationsHandler(c *gin.Context) {
	// Access the database connection
	db := services.GetConnection()

	// Declare a slice to store all reservations
	var reservations []models.Reservation

	// Perform a query to retrieve all reservations from the database
	if err := db.Preload("User").Preload("HairSalon").Preload("HairSalon").Find(&reservations).Error; err != nil {
		// In case of an error during the query, return an error response
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reservations"})
		return
	}

	// Respond with the retrieved reservations
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
