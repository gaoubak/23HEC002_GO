package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"

	"github.com/gin-gonic/gin"
)

// DeleteHairDresserHandler handles the deletion of a HairDresser
func DeleteHairDresserHandler(c *gin.Context) {
	// Retrieve the ID of the HairDresser from the URL parameters
	HairDresserID := c.Param("HairDresserId")

	// Access the database connection
	db := services.GetConnection()

	// Retrieve the HairDresser from the database using its ID
	var existingHairDresser models.HairDresser
	if err := db.First(&existingHairDresser, HairDresserID).Error; err != nil {
		// Display the error and return a NotFound response
		fmt.Println("Error retrieving HairDresser from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "HairDresser not found"})
		return
	}

	// Display information for debugging
	fmt.Println("Existing HairDresser to delete:", existingHairDresser)

	// Delete the HairDresser from the database
	if err := db.Delete(&existingHairDresser).Error; err != nil {
		// Display the error and return an InternalServerError response
		fmt.Println("Error deleting HairDresser from the database:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete HairDresser"})
		return
	}

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"message": "HairDresser deleted successfully"})
}
