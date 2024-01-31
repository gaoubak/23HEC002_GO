package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"
	validators "GolandProject/validators/reservation"

	"github.com/gin-gonic/gin"
)

// UpdateReservationHandler handles the update of a reservation
func UpdateReservationHandler(c *gin.Context) {
	// Retrieve the ID of the reservation from the URL parameters
	reservationID := c.Param("reservationId")

	// Access the database connection
	db := services.GetConnection()

	// Retrieve the existing reservation from the database using its ID
	var existingReservation models.Reservation
	if err := db.First(&existingReservation, reservationID).Error; err != nil {
		// Display the error and return a NotFound response
		fmt.Println("Error retrieving reservation from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	// Display information for debugging
	fmt.Println("Existing Reservation:", existingReservation)

	// Create a new UpdateReservationValidator and bind JSON data
	var reservationValidator validators.UpdateReservationValidator
	if err := c.ShouldBindJSON(&reservationValidator); err != nil {
		// Display the validation error and return a BadRequest response
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the UpdateReservationValidator structure
	if err := reservationValidator.Validate(); err != nil {
		// Display the validation error and return a BadRequest response
		fmt.Println("Error validating reservation input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the existing reservation with the validated data
	existingReservation.Update(reservationValidator)

	// Save the updated reservation to the database
	db.Save(&existingReservation)

	// Respond with the updated reservation data
	c.JSON(http.StatusOK, gin.H{"reservation": existingReservation})
}
