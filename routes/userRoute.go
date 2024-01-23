// router/userRoute.go
package routes

import (
	"GolandProject/contexts"
	"GolandProject/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Groupe de routes avec le middleware UserContext pour les routes sous /user
	userGroup := router.Group("/user")
	userGroup.Use(contexts.UserContext())

	// Définissez vos gestionnaires qui peuvent accéder à l'utilisateur depuis le contexte
	userGroup.GET("/:userId", handlers.UserHandler)
	userGroup.PUT("/:userId", handlers.UpdateUserHandler)
}
