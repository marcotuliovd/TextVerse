package main

import (
	"api/src/config"
	"api/src/database"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Loading()
	database.DBConnection()
	fmt.Println("rodando api")

	r  := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}