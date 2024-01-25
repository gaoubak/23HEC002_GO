package handlers

import (
	"GolandProject/services"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const userKey = "user"

// LoginForm est la structure pour le formulaire de connexion
type LoginForm struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func LoginHandler(c *gin.Context) {
	session := sessions.Default(c)

	var loginForm LoginForm
	if err := c.ShouldBindJSON(&loginForm); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validez les données du formulaire
	if strings.Trim(loginForm.Username, " ") == "" || strings.Trim(loginForm.Password, " ") == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Les paramètres ne peuvent pas être vides"})
		return
	}

	// Vérifiez la correspondance du nom d'utilisateur et du mot de passe (par exemple, à partir de la base de données)
	if isValidUser(loginForm.Username, loginForm.Password) {
		// Enregistrez le nom d'utilisateur dans la session
		session.Set(userKey, loginForm.Username)
		session.Save()

		c.JSON(http.StatusOK, gin.H{"message": "Authentification réussie"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Échec de l'authentification"})
	}
}

// LogoutHandler est le gestionnaire pour la déconnexion des utilisateurs
func LogoutHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete(userKey)
	session.Save()

	c.JSON(http.StatusOK, gin.H{"message": "Déconnexion réussie"})
}

func isValidUser(username, password string) bool {
	// Vérifiez si le nom d'utilisateur et le mot de passe correspondent à un utilisateur enregistré dans la base de données
	user, err := services.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return false
	}
	return user != nil
}
