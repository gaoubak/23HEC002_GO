package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// GetReservationsByClientIDHandler retrieves all reservations by ClientID
func GetReservationsByClientIDHandler(c *gin.Context) {
	// Retrieve the ClientID from the URL parameters
	clientID := c.Param("clientId")

	// Access the database connection
	db := services.GetConnection()

	// Declare a slice to store all reservations
	var reservations []models.Reservation

	// Perform a query to retrieve all reservations with the given ClientID from the database
	if err := db.Where("client_id = ?", clientID).Find(&reservations).Error; err != nil {
		// Display the error and return a InternalServerError response
		fmt.Println("Error retrieving reservations by ClientID from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve reservations"})
		return
	}

	// Respond with the retrieved reservations
	c.JSON(http.StatusOK, gin.H{"reservations": reservations})
}
