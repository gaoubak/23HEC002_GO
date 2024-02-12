package handlers

import (
	"GolandProject/email"
	"GolandProject/models"
	"GolandProject/services"
	validators "GolandProject/validators/reservation"
	"fmt"
	"net/http"
	"strconv"

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
		HairDresserID:     reservation.HairDresserID,
		UserID:            reservation.UserID,
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
	if err := db.Preload("User").Create(&reservation).Error; err != nil {
		fmt.Println("Error creating reservation:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}

	// Send confirmation email
	err := sendConfirmationEmail(reservation)
	if err != nil {
		fmt.Println("Error sending confirmation email:", err)
	}

	c.JSON(http.StatusCreated, gin.H{"reservation": reservation})
}

func sendConfirmationEmail(reservation models.Reservation) error {
	// Construct the email message
	to := reservation.User.Email
	subject := "Reservation Confirmation"
	body := "Your reservation has been successfully created."

	// You need to provide your SMTP server configuration
	smtpServer := "smtp.gmail.com"
	smtpPort := 465
	smtpUser := "kader"
	smtpPassword := "nuax janz biav hrfn"

	// Convert smtpPort to string
	portStr := strconv.Itoa(smtpPort)

	// Send the email
	err := email.SendConfirmationEmail(to, subject, body, smtpServer, portStr, smtpUser, smtpPassword)
	if err != nil {
		return err
	}

	return nil
}
