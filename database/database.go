package database

import (
	"animal-rescue-be/helpers"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

var (
	DB_HOST     = ""
	DB_PORT     = ""
	DB_USER     = ""
	DB_PASSWORD = ""
	DB_NAME     = ""
)

func initEnvVariables() {
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
}

func InitDatabase() {
	initEnvVariables()
	psqlConnInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	database, err := gorm.Open("postgres", psqlConnInfo)
	helpers.HandleErr(err)

	//change to use global variables
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)

	DB = database
}
