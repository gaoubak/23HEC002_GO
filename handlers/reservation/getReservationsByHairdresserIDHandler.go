package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// GetReservationsByHairdresserIDHandler retrieves all reservations by HairdresserID
func GetReservationsByHairdresserIDHandler(c *gin.Context) {
	// Retrieve the HairdresserID from the URL parameters
	hairdresserID := c.Param("hairdresserId")

	// Access the database connection
	db := services.GetConnection()

	// Declare a slice to store all reservations
	var reservations []models.Reservation

	// Perform a query to retrieve all reservations with the given HairdresserID from the database
	if err := db.Where("hairdresser_id = ?", hairdresserID).Find(&reservations).Error; err != nil {
		// Display the error and return a InternalServerError response
		fmt.Println("Error retrieving reservations by HairdresserID from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reservations"})
		return
	}

	// Respond with the retrieved reservations
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
