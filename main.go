package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", GetUser).Methods("GET")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/user", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}

var db *gorm.DB
var err error

func initializeMigration() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	name := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")
	fmt.Println(host, name, port)
	dbURI := fmt.Sprintf("host=localhost user=%s password=%s dbname=go_rest_api port=5432 sslmode=disable TimeZone=Asia/Kolkata", user, password)

	db, err = gorm.Open(postgres.Open(dbURI), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to the Database")
	}
	db.AutoMigrate(&User{})
}

func main() {
	initializeMigration()
	initializeRouter()
}
