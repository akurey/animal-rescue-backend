package main

import (
	"animal-rescue-be/controllers"
	middleware "animal-rescue-be/middleware"
	routes "animal-rescue-be/routes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db = make(map[string]string)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}
	// gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())
	router.GET("/test", func(ctx *gin.Context) { // Returns success if the request includes a valid json web token
		ctx.JSON(http.StatusOK, gin.H{"success": true, "data": nil, "message": "already logged in"})
	})

	ping := new(controllers.PingController)
	// Ping test
	router.GET("/ping", ping.Ping)

	router.Run(":" + port)
}
