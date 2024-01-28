package handlers

import (
	"GolandProject/models"
	"GolandProject/services"
	validators "GolandProject/validators/user"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUserHandler(c *gin.Context) {
	var user models.User

	// Bind the JSON request body to the user model
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate the user data for registration
	userValidator := validators.RegisterUserValidator{
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Password:  user.Password,
	}

	if err := userValidator.Validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create the user using the model method
	user.Create(userValidator)

	// Access the database connection
	db := services.GetConnection()

	// Hash the user's password before saving to the database
	hashedPassword, err := user.HashPassword()
	if err != nil {
		fmt.Println("Error hashing password:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}
	user.EncryptedPassword = string(hashedPassword)

	// Save the new user to the database
	if err := db.Create(&user).Error; err != nil {
		fmt.Println("Error creating user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	user.Password = "" // Do not expose the password in the response
	c.JSON(http.StatusCreated, gin.H{"user": user})
}
