// router/userRoute.go
package routes

import (
	// "GolandProject/contexts"
	"GolandProject/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Utilisez le middleware UserContext pour extraire l'ID de l'URL et ajouter l'utilisateur au contexte
	// router.Use(contexts.UserContext())

	// Définissez vos gestionnaires qui peuvent accéder à l'utilisateur depuis le contexte
	router.GET("/user/:userId", handlers.UserHandler)
}
