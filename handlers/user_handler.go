package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BookAppointment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Appointment booked successfully",
	})
}

func UserGetAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Get User successfully",
	})
}
