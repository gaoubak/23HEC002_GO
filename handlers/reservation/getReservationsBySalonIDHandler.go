package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// GetReservationsBySalonIDHandler retrieves all reservations by SalonID
func GetReservationsBySalonIDHandler(c *gin.Context) {
	// Retrieve the SalonID from the URL parameters
	salonID := c.Param("salonId")

	// Access the database connection
	db := services.GetConnection()

	// Declare a slice to store all reservations
	var reservations []models.Reservation

	// Perform a query to retrieve all reservations with the given SalonID from the database
	if err := db.Preload("Client").Where("salon_id = ?", salonID).Find(&reservations).Error; err != nil {
		// Display the error and return a InternalServerError response
		fmt.Println("Error retrieving reservations by SalonID from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reservations"})
		return
	}

	// Respond with the retrieved reservations
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
