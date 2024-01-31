package contexts

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const reservationKey = "reservation"

func ReservationContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		reservationID := c.Param("reservationId")

		var reservation models.Reservation

		db := services.GetConnection()

		db.First(&reservation, reservationID)

		if reservation.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
			c.Abort()
			return
		}

		c.Set(reservationKey, &reservation)

		c.Next()
	}
}
