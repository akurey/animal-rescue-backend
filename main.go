package main

import (
	"animal-rescue-be/controllers"
	"animal-rescue-be/database"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	engine := gin.Default()

	animal := new(controllers.AnimalController)

	form := new(controllers.FormController)

	health := new(controllers.HealthController)

	report := new(controllers.ReportController)

	engine.GET("/", health.Check)

	engine.GET("/animals", animal.GetAnimals)

	engine.GET("/form/:id/fields", form.GetFormFields)

	engine.GET("/reports", report.GetReports)

	engine.POST("/reports", report.AddReport)

	engine.PATCH("/reports", report.UpdateReport)

	engine.DELETE("/reports", report.DeleteReport)

	return engine
}

func init() {
	if os.Getenv("APP_ENV") != "production" {
		if godotenv.Load(".env") != nil {
			log.Fatal("Error loading .env file")
		}
	}
}

func main() {
	database.InitDatabase()
	router := setupRouter()
	port := "8080"

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	// Listen and Server in 0.0.0.0:8080
	router.Run(fmt.Sprintf(":%v", port))
}
