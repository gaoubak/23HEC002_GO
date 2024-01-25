// contexts/userContext.go
package contexts

import (
	"GolandProject/models"
	"GolandProject/services"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userKey = "user"

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

		c.Set(userKey, &user)

		// Vérifie si l'utilisateur est authentifié et l'enregistre dans la session
		session := sessions.Default(c)
		if authUser, exists := c.Get("authenticatedUser"); exists {
			session.Set(userKey, authUser)
			session.Save()
		}

		c.Next()
	}
}
