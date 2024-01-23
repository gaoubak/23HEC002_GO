package main

import (
	"GolandProject/models"
	"GolandProject/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Starting server on port 8080...")

	// Initialize the database connection
	database, err := services.InitSqlConnection()
	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}

	// Get a reference to the underlying database connection
	connection, err := database.DB()
	if err != nil {
		log.Fatal("An error occurred with the database connection: ", err)
	}
	defer connection.Close()

	// Perform database migrations for models
	err = database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Error performing database migrations: ", err)
	}

	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.GET("/reste", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ponggararararara",
		})
	})

	router.Run(":8080")
}
