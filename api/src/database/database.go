package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DNS = ""
var DB *gorm.DB

func DBConnection() {
	var error error
	DNS = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", os.Getenv("HOST"), os.Getenv("USER_DATABASE"), os.Getenv("PASSWORD_DATABASE"), os.Getenv("NAME_DATABASE"), os.Getenv("PORT"))
	DB, error = gorm.Open(postgres.Open(DNS), &gorm.Config{})
	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB Connection")
	}
}