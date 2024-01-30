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

	// Retrieve the reservation from the database using its ID
	var existingReservation models.Reservation
	if err := db.First(&existingReservation, reservationID).Error; err != nil {
		// Display the error and return a NotFound response
		fmt.Println("Error retrieving reservation from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}

	// Display information for debugging
	fmt.Println("Existing Reservation:", existingReservation)

	// Access the reservation from the context
	reservation, exists := c.Get("reservation")
	if !exists {
		// Display an error message
		fmt.Println("Reservation information not available for this request put")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Reservation information not available for this request put"})
		return
	}

	// Ensure to convert the reservation to type *models.Reservation
	updateReservation, ok := reservation.(*models.Reservation)
	if !ok {
		// Display an error message
		fmt.Println("Reservation information not available in the expected format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Reservation information not available in the expected format"})
		return
	}

	var reservationValidator validators.UpdateReservationValidator

	// Read and validate the JSON data sent by the user
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

	// Perform the update of the reservation
	updateReservation.Update(reservationValidator)

	// Save the updated reservation to the database
	db.Save(updateReservation)

	// Respond with the updated reservation data
	c.JSON(http.StatusOK, gin.H{"reservation": updateReservation})
}
