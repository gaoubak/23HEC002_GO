package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// GetReservationsByUserIDHandler retrieves all reservations by UserID
func GetReservationsByUserIDHandler(c *gin.Context) {
	// Retrieve the UserID from the URL parameters
	UserID := c.Param("UserID")

	// Access the database connection
	db := services.GetConnection()

	// Declare a slice to store all reservations
	var reservations []models.Reservation

	// Perform a query to retrieve all reservations with the given UserID from the database
	if err := db.Preload("User").Preload("HairSalon").Where("user_id = ?", UserID).Find(&reservations).Error; err != nil {
		// Display the error and return a InternalServerError response
		fmt.Println("Error retrieving reservations by UserID from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reservations"})
		return
	}

	// Respond with the retrieved reservations
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
