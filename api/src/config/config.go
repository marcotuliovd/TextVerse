package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Connection = ""
	Port = 0
)

func Loading() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 5000
	}

	Connection = fmt.Sprintf("%s:%s@/%s",
		os.Getenv("USER"),
		os.Getenv("PASSWORD_DATABASE"),
		os.Getenv("NAME_DATABASE"),
	)
}