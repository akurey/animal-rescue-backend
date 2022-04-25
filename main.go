package main

import (
	"animal-rescue-be/controllers"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	engine := gin.Default()

	ping := new(controllers.PingController)

	// Ping test
	engine.GET("/ping", ping.Ping)

	return engine
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
