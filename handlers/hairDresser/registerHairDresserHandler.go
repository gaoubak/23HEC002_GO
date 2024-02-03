package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	validators "GolandProject/validators/hairDresser"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RegisterHairDresserHandler handles the registration of a new HairDresser
func RegisterHairDresserHandler(c *gin.Context) {
	var HairDresser models.HairDresser

	// Bind the JSON request body to the HairDresser model
	if err := c.ShouldBindJSON(&HairDresser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the HairDresser data
	HairDresserValidator := validators.CreateHairDresserValidator{
		Name:        HairDresser.Name,
		Email:       HairDresser.Email,
		Speciality:  HairDresser.Speciality,
		Description: HairDresser.Description,
		//ReservationID: HairDresser.ReservationID,
	}

	if err := HairDresserValidator.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the HairDresser using the model method
	HairDresser.Create(HairDresserValidator)

	// Access the database connection
	db := services.GetConnection()

	// Save the new HairDresser to the database
	if err := db.Create(&HairDresser).Error; err != nil {
		fmt.Println("Error creating HairDresser:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create HairDresser"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"HairDresser": HairDresser})
}
