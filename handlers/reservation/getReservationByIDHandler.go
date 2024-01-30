package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// GetReservationByIDHandler retrieves one reservation by ID
func GetReservationByIDHandler(c *gin.Context) {
	// Retrieve the ID of the reservation from the URL parameters
	reservationID := c.Param("reservationId")

	// Access the database connection
	db := services.GetConnection()

	// Retrieve the reservation from the database using its ID
	var existingReservation models.Reservation
	if err := db.First(&existingReservation, reservationID).Error; err != nil {
		// Display the error and return a NotFound response
		fmt.Println("Error retrieving reservation from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	// Respond with the retrieved reservation
	c.JSON(http.StatusOK, gin.H{"reservation": existingReservation})
}
