package controllers

import (
	"api/src/database"
	modelos "api/src/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		log.Fatal(erro)
	}

	var user modelos.Usuario
	if erro = json.Unmarshal(bodyRequest, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro = database.DBConnection()
	if erro != nil {
		log.Fatal(erro)
	}
}

func FindUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("busca todos"))
}

func FindUserbyId(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("busca por id"))
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("atualiza user"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criand user"))
}