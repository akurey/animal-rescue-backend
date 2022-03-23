package main

import (
	"github.com/gin-gonic/gin"

	"animal-rescue-be/controllers"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	ping := new(controllers.PingController)

	// Ping test
	r.GET("/ping", ping.Ping)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
