package main

import (
	"GolandProject/handlers"
	"GolandProject/models"
	"GolandProject/routes"
	"GolandProject/seeders"
	"GolandProject/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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
	err = database.AutoMigrate(&models.User{}, &models.HairSalon{}, &models.Reservation{}, &models.HairDresser{})
	if err != nil {
		log.Fatal("Error performing database migrations: ", err)
	}

	// Seed the database with some initial data
	seeders.HairSalonSeeder(database)
	seeders.UserSeeder(database)
	if err != nil {
		log.Fatal("Error seeding the database: ", err)
	}

	// Initialisez le routeur Gin
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		
		c.Next()
	})
	
	// Setup the cookie store for session management
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// Routes d'authentification
	router.POST("/register", handlers.RegisterUserHandler)
	router.POST("/login", handlers.LoginHandler)
	router.GET("/logout", handlers.LogoutHandler)

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
