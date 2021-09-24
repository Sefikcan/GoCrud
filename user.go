package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "ttt:ppp@tcp(127.0.0.1:3306)/db?parseTime=true"

type User struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}

	DB.AutoMigrate(&User{})
}

// Get User List Operation
func GetUsers(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var users []User
	DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}

// Get Single User
func GetUser(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var user User
	params := mux.Vars(r)

	DB.First(&user, params["id"])
	json.NewEncoder(w).Encode(user)
}

// Create User Operation
func CreateUser(w http.ResponseWriter, r *http.Request) {
	setHeader(w)
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

// Update User Operation
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	var user User
	params := mux.Vars(r)

	DB.First(&user, params["id"])
	json.NewDecoder(r.Body).Decode(&user)
	DB.Save(&user)

	json.NewEncoder(w).Encode(user)
}

// Delete User Operation
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	setHeader(w)

	var user User
	params := mux.Vars(r)

	DB.Delete(&user, params["id"])
	json.NewEncoder(w).Encode("user is deleted successfully")
}

func setHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
