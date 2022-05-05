package repositories

import (
	models "animal-rescue-be/models"
)

var currentIndex int
var users [100]models.User

// this file will be updated based on the database/orm implementation.

func GetUserByEmail(email string) models.User {
	var user models.User
	password := "password"
	user.Email = &email
	user.Password = &password
	return user
}

func UserExists(email string) bool {
	// Check if the user exists
	return false
}

func AddUser(user models.User) {
	users[currentIndex] = user
	currentIndex = currentIndex + 1
}
