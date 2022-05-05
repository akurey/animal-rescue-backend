package controllers

import (
	"animal-rescue-be/models"
	repositories "animal-rescue-be/repositories"
	utils "animal-rescue-be/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

//sign up user
func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		validationErr := validate.Struct(user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "data": nil, "message": validationErr.Error()})
			return
		}
		userExists := repositories.UserExists(*user.Email)
		password := utils.HashPassword(*user.Password)
		user.Password = &password
		if userExists {
			c.JSON(http.StatusInternalServerError,
				gin.H{"success": false, "data": nil, "message": "email ready exists"})
			return
		}
		user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		user.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		token, refreshToken, _ := utils.GenerateAllTokens(*user.Email)
		user.Token = &token
		user.Refresh_token = &refreshToken
		repositories.AddUser(user)
		c.JSON(http.StatusOK, gin.H{"success": true, "data": nil, "message": "user signup sucess"})
	}
}

//Login user
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		var foundUser models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		foundUser = repositories.GetUserByEmail(*user.Email)
		passwordIsValid, msg := utils.VerifyPassword(*user.Password, *foundUser.Password)
		if passwordIsValid != true {
			c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
			return
		}
		if foundUser.Email == nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "user not found"})
			return
		}
		token, refreshToken, _ := utils.GenerateAllTokens(*foundUser.Email)
		c.JSON(http.StatusOK, gin.H{"success": true, "data": gin.H{
			"token":         token,
			"refresh_token": refreshToken},
			"message": "success"})
	}
}
