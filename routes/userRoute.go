// router/userRoute.go
package routes

import (
	userHandler "GolandProject/handlers/user"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/user/:userId", userHandler.UserHandler)
		userGroup.POST("/register", userHandler.RegisterUserHandler)
	}
	// Define your handlers that can access the user from the context
}
