package repositories

import (
	models "animal-rescue-be/models"
)

var currentIndex int
var users [100]models.User

func GetUserByEmail(email string) models.User {
	var user models.User
	e := "email@gmail.com"
	user.Email = &e
	return user
}

func UserExists(email string) bool {
	return false
}

func AddUser(user models.User) {
	users[currentIndex] = user
	currentIndex = currentIndex + 1
}
