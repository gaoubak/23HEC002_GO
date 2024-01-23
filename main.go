package main

import (
	"GolandProject/models"
	"GolandProject/routes"
	"GolandProject/services"
	"fmt"
	"log"
	"net/http"

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

	// Initialisez le routeur Gin
	router := gin.Default()

	// Configurez les routes en utilisant le package routeurs
	routes.SetupRoutes(router)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// Lancez votre application
	router.Run(":8000")
}
