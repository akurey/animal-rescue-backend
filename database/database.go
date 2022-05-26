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
	host     = os.Getenv("HOST")
	port     = os.Getenv("PORT")
	user     = os.Getenv("USR")
	password = os.Getenv("PASS")
	dbname   = os.Getenv("DATABASE")
)

func InitDatabase() {
	psqlConnInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	database, err := gorm.Open("postgres", psqlConnInfo)
	helpers.HandleErr(err)

	//change to use global variables
	database.DB().SetMaxIdleConns(20)
	database.DB().SetMaxOpenConns(200)

	DB = database
}
