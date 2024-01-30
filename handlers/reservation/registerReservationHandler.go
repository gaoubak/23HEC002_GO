package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	validators "GolandProject/validators/reservation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterReservationHandler handles the registration of a new reservation
func RegisterReservationHandler(c *gin.Context) {
	var reservation models.Reservation

	// Bind the JSON request body to the reservation model
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status for the reservation
	reservation.Status = "En Attente"

	// Validate the reservation data
	reservationValidator := validators.CreateReservationValidator{
		DateOfReservation: reservation.DateOfReservation,
		Status:            reservation.Status,
		SalonID:           reservation.SalonID,
		HairdresserID:     reservation.HairdresserID,
		ClientID:          reservation.ClientID,
	}

	if err := reservationValidator.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the reservation using the model method
	reservation.Create(reservationValidator)

	// Access the database connection
	db := services.GetConnection()

	// Save the new reservation to the database
	if err := db.Create(&reservation).Error; err != nil {
		fmt.Println("Error creating reservation:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"reservation": reservation})
}
