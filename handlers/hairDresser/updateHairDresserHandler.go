package handlers

import (
	"fmt"
	"net/http"

	"GolandProject/models"
	"GolandProject/services"
	validators "GolandProject/validators/hairDresser"

	"github.com/gin-gonic/gin"
)

// UpdateHairDresserHandler handles the update of a HairDresser
func UpdateHairDresserHandler(c *gin.Context) {
	// Retrieve the ID of the HairDresser from the URL parameters
	HairDresserID := c.Param("HairDresserId")

	// Access the database connection
	db := services.GetConnection()

	// Retrieve the existing HairDresser from the database using its ID
	var existingHairDresser models.HairDresser
	if err := db.First(&existingHairDresser, HairDresserID).Error; err != nil {
		// Display the error and return a NotFound response
		fmt.Println("Error retrieving HairDresser from the database:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "HairDresser not found"})
		return
	}

	// Display information for debugging
	fmt.Println("Existing HairDresser:", existingHairDresser)

	// Create a new UpdateHairDresserValidator and bind JSON data
	var HairDresserValidator validators.UpdateHairDresserValidator
	if err := c.ShouldBindJSON(&HairDresserValidator); err != nil {
		// Display the validation error and return a BadRequest response
		fmt.Println("Error binding JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the UpdateHairDresserValidator structure
	if err := HairDresserValidator.Validate(); err != nil {
		// Display the validation error and return a BadRequest response
		fmt.Println("Error validating HairDresser input:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the existing HairDresser with the validated data
	existingHairDresser.Update(HairDresserValidator)

	// Save the updated HairDresser to the database
	db.Save(&existingHairDresser)

	// Respond with the updated HairDresser data
	c.JSON(http.StatusOK, gin.H{"HairDresser": existingHairDresser})
}
