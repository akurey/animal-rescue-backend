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

func InitServerHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin,X-Requested-With,Content-Type,Accept,content-type,application/json")
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()

	engine := gin.Default()

	engine.Use(InitServerHeaders())

	animal := new(controllers.AnimalController)

	form := new(controllers.FormController)

	health := new(controllers.HealthController)

	report := new(controllers.ReportController)

	user := new(controllers.UserController)

	engine.GET("/", health.Check)

	engine.GET("/animals", animal.GetAnimals)

	engine.GET("/form/:id/fields", form.GetFormFields)
	
	engine.GET("/report/:id", report.GetAnimalRecord)

	engine.GET("/form/address", form.GetAddressOptions)

	engine.GET("/reports", report.GetReports)

	engine.POST("/reports", report.AddReport)

	engine.PATCH("/reports/:id", report.UpdateReport)

	engine.DELETE("/reports/:id", report.DeleteReport)

	engine.POST("/users", user.SignUpUser)

	engine.POST("/users/login", user.LoginUser)

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
