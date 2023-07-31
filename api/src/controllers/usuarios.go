package controllers

import "net/http"

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("criando user"))
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