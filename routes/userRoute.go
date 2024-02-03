// router/userRoute.go
package routes

import (
	"GolandProject/contexts"
	hairDresserHandler "GolandProject/handlers/hairDresser"
	hairSalonHandler "GolandProject/handlers/hairSalon"
	reservationHandler "GolandProject/handlers/reservation"
	userHandler "GolandProject/handlers/user"
	"GolandProject/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Define your handlers that can access the user from the context
	// Groupe de routes avec le middleware UserContext pour les routes sous /user
	userGroup := router.Group("/user")
	reservationGroup := router.Group("/reservation")
	hairSalonGroup := router.Group("/hair-salon")
	hairDresserGroup := router.Group("/hair-dresser")
	// Appliquez AuthRequired aux routes nécessitant une authentification
	userGroup.Use(middleware.AuthRequired)
	hairSalonGroup.Use(middleware.AuthRequired)
	//hairDresserGroup.Use(middleware.AuthRequired)
	// Appliquez le middleware UserContext aux routes sous /user
	userGroup.Use(contexts.UserContext())
	userGroup.Use(contexts.ReservationContext())
	hairDresserGroup.Use(contexts.HairDresserContext())

	{
		hairDresserGroup.PUT("/:hairdresserId", hairDresserHandler.UpdateHairDresserHandler)
		hairDresserGroup.GET("/", hairDresserHandler.GetAllHairDresserHandler)
		hairDresserGroup.DELETE("/:hairdresserId", hairDresserHandler.DeleteHairDresserHandler)
		hairDresserGroup.POST("/", hairDresserHandler.RegisterHairDresserHandler)
		hairDresserGroup.GET("/:hairdresserId", hairDresserHandler.GetSingleHairDresserHandler)
	}
	{
		// Définissez vos gestionnaires qui peuvent accéder à l'utilisateur depuis le contexte
		userGroup.GET("/:userId", userHandler.UserHandler)
		userGroup.PUT("/:userId", userHandler.UpdateUserHandler)
		userGroup.GET("/", userHandler.GetAllUserHandler)
	}

	{
		reservationGroup.GET("/", reservationHandler.GetAllReservationsHandler)
		reservationGroup.GET("/:reservationId", reservationHandler.GetReservationByIDHandler)
		reservationGroup.GET("/User/:userID", reservationHandler.GetReservationsByUserIDHandler)
		reservationGroup.GET("/hairdresser/:hairdresserId", reservationHandler.GetReservationsByHairdresserIDHandler)
		reservationGroup.GET("/salon/:salonId", reservationHandler.GetReservationsBySalonIDHandler)
		reservationGroup.DELETE("/:reservationId", reservationHandler.DeleteReservationHandler)
		reservationGroup.PUT("/:reservationId", reservationHandler.UpdateReservationHandler)
		reservationGroup.POST("/", reservationHandler.RegisterReservationHandler)
	}
	{
		hairSalonGroup.GET("/:hairSalonId", hairSalonHandler.GetSingleHairSalonHandler)
		hairSalonGroup.GET("/", hairSalonHandler.GetAllHairSalonHandler)
	}
}
