package main

import (
	userHandler "gocrud/handlers/user"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/api/v1/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/api/v1/users/{id}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/api/v1/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/users/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/api/v1/users/{id}", userHandler.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8081", r))
}

func main() {
	userHandler.InitialMigration()
	initializeRouter()
}
