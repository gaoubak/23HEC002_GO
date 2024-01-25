// middleware/middleware.go
package middleware

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userkey = "user"

// AuthRequired est un middleware pour vérifier la session
func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		// L'utilisateur n'est pas authentifié
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}
	// L'utilisateur est authentifié
	c.Next()
}
