package routes

import (
	"GolandProject/handlers"

	"github.com/gin-gonic/gin"
)

func SetUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/book-appointment", handlers.BookAppointment)
		userGroup.GET("", handlers.UserGetAll)
	}
}
