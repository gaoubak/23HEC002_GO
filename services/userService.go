// services/userService.go
package services

import (
	"GolandProject/models"
	"fmt"
)

// GetUserByUsernameAndPassword récupère un utilisateur à partir de la base de données en fonction du nom d'utilisateur et du mot de passe
func GetUserByUsernameAndPassword(username, password string) (*models.User, error) {
	var user models.User
	fmt.Println("usernameCaca: ", username)
	if err := GetConnection().Where("username = ? AND password = ?", username, password).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
