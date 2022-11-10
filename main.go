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

	animal := new(controllers.AnimalController)

	form := new(controllers.FormController)

	health := new(controllers.HealthController)

	report := new(controllers.ReportController)

	engine.GET("/", health.Check)

	engine.GET("/animals", animal.GetAnimals)

	engine.GET("/form/:id/fields", form.GetFormFields)
	
	engine.GET("/report/:id", report.GetAnimalRecord)

	engine.GET("/form/address", form.GetAddressOptions)

	engine.GET("/reports", report.GetReports)

	engine.POST("/reports", report.AddReport)

	engine.PATCH("/reports/:id", report.UpdateReport)

	engine.DELETE("/reports/:id", report.DeleteReport)

	return engine
}

func init() {
	if godotenv.Load(".env") != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	database.InitDatabase()
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
