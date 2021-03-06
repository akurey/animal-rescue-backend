package database

import (
	"animal-rescue-be/helpers"
	"fmt"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var DB *gorm.DB

var (
	DB_HOST           = ""
	DB_PORT           = ""
	DB_USER           = ""
	DB_PASSWORD       = ""
	DB_NAME           = ""
	DB_MAX_IDLE_CONNS = ""
	DB_MAX_OPEN_CONNS = ""
)

func initEnvVariables() {
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_MAX_IDLE_CONNS = os.Getenv("DB_MAX_IDLE_CONNS")
	DB_MAX_OPEN_CONNS = os.Getenv("DB_MAX_OPEN_CONNS")
}

func InitDatabase() {
	initEnvVariables()
	psqlConnInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", DB_HOST, DB_PORT, DB_USER, DB_PASSWORD, DB_NAME)
	database, err := gorm.Open("postgres", psqlConnInfo)
	helpers.HandleErr(err)

	maxIddleConns, err := strconv.Atoi(DB_MAX_IDLE_CONNS)
	if err != nil {
		helpers.HandleErr(err)
	} else {
		database.DB().SetMaxIdleConns(maxIddleConns)
	}

	maxOpenConns, err := strconv.Atoi(DB_MAX_OPEN_CONNS)
	if err != nil {
		helpers.HandleErr(err)
	} else {
		database.DB().SetMaxOpenConns(maxOpenConns)
	}

	DB = database
}
