package contexts

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

const hairDresserKey = "hairDresser"

func HairDresserContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		hairDresserID := c.Param("hairDresserId")

		var hairDresser models.HairDresser

		db := services.GetConnection()

		db.First(&hairDresser, hairDresserID)

		if hairDresser.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "hairDresser not found"})
			c.Abort()
			return
		}

		c.Set(hairDresserKey, &hairDresser)

		c.Next()
	}
}
