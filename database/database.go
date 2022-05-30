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
	DB_HOST     = os.Getenv("HOST")
	DB_PORT     = os.Getenv("PORT")
	DB_USER     = os.Getenv("USR")
	DB_PASSWORD = os.Getenv("PASS")
	DB_NAME     = os.Getenv("DATABASE")
)

func InitDatabase() {
	psqlConnInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	database, err := gorm.Open("postgres", psqlConnInfo)
	helpers.HandleErr(err)

	//change to use global variables
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)

	DB = database
}
