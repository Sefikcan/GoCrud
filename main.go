package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}
