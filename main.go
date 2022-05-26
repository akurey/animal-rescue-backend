package main

import (
	"animal-rescue-be/controllers"
	"animal-rescue-be/database"

	"log"

	"github.com/joho/godotenv"

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

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")

	database.InitDatabase()
}
