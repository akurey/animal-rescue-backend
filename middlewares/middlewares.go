package middlewares

import (
	"net/http"

	"animal-rescue-be/database"
	"animal-rescue-be/models"
	"github.com/gin-gonic/gin"
	"animal-rescue-be/helpers"
)

const unauthorizedMessage = "Authentication credentials were not provided."
const unauthorizedMessageCode = "not_authenticated"

type Token struct {
    token  string `gorm:"not null"`
}

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		errorResponse := models.ErrorResponse{
			Error: struct {
				Message string `json:"message"`
				Code    string `json:"code"`
			}{
				Message: unauthorizedMessage,
				Code:   unauthorizedMessageCode,
			},
			Code: -1,
		}

		err := helpers.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		tokenString := helpers.ExtractToken(c)
		var validToken string
		result := database.DB.Table("\"AP_Users\" AU").
				  Select("AU.token").
				  Where("AU.token = ?", tokenString).
				  Find(&validToken).Error

		if result != nil || validToken == "" {
			c.JSON(http.StatusUnauthorized, errorResponse)
			c.Abort()
			return
		}

		c.Next()
	}
}