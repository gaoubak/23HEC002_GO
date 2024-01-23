// contexts/userContext.go
package contexts

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Param("userId")

		var user models.User

		db := services.GetConnection()

		db.First(&user, userID)

		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			c.Abort()
			return
		}

		c.Set("user", &user)

		c.Next()
	}
}
