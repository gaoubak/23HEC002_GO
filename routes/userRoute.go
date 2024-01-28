// router/userRoute.go
package routes

import (
	"GolandProject/contexts"
	userHandler "GolandProject/handlers/user"
	"GolandProject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define your handlers that can access the user from the context
	// Groupe de routes avec le middleware UserContext pour les routes sous /user
	userGroup := router.Group("/user")
	userGroup.Use(contexts.UserContext())

	// Appliquez AuthRequired aux routes nécessitant une authentification
	userGroup.Use(middleware.AuthRequired)

	{
		// Définissez vos gestionnaires qui peuvent accéder à l'utilisateur depuis le contexte
		userGroup.GET("/:userId", userHandler.UserHandler)
		userGroup.PUT("/:userId", userHandler.UpdateUserHandler)
	}
}
