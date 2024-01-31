package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// DeleteReservationHandler handles the deletion of a reservation
func DeleteReservationHandler(c *gin.Context) {
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

	// Display information for debugging
	fmt.Println("Existing Reservation to delete:", existingReservation)

	// Delete the reservation from the database
	if err := db.Delete(&existingReservation).Error; err != nil {
		// Display the error and return an InternalServerError response
		fmt.Println("Error deleting reservation from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reservation"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
}
