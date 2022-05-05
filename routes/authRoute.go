package routes

import (
	controller "animal-rescue-be/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.POST("/users/signup", controller.SignUp())
	route.POST("/users/login", controller.Login())
}
