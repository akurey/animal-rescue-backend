package main

import (
	"animal-rescue-be/controllers"
	"animal-rescue-be/database"
	"animal-rescue-be/middlewares"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func InitServerHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
	}
}

func setupRouter() *gin.Engine {

	engine := gin.Default()
	engine.Use(InitServerHeaders())

	apiPublic := engine.Group("/api")
	apiProtected := engine.Group("/api")
	apiProtected.Use(middlewares.JwtAuthMiddleware())
	
	animal := new(controllers.AnimalController)
	form := new(controllers.FormController)
	health := new(controllers.HealthController)
	report := new(controllers.ReportController)
	user := new(controllers.UserController)

	// apiProtected.OPTIONS("/*path", func(c *gin.Context) {
	// 	c.Header("Access-Control-Allow-Methods", "POST")
	// 	c.Header("Access-Control-Allow-Headers", "content-type")
	// })

	apiProtected.GET("/animals", animal.GetAnimals)
	apiProtected.GET("/form/:id/fields", form.GetFormFields)
	apiProtected.GET("/report/:id", report.GetAnimalRecord)
	apiProtected.GET("/form/address", form.GetAddressOptions)
	apiProtected.GET("/reports", report.GetReports)
	apiProtected.POST("/reports", report.AddReport)
	apiProtected.PATCH("/reports/:id", report.UpdateReport)
	apiProtected.DELETE("/reports/:id", report.DeleteReport)
	apiProtected.POST("/users/logout", user.LogoutUser)

	apiPublic.GET("/", health.Check)
	apiPublic.POST("/users", user.SignUpUser)
	apiPublic.POST("/users/login", user.LoginUser)

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
