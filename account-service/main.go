package main

import (
	"account-service/handlers"
	"account-service/models"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {

	db := openDbConnection()

	createDbModels(db)

	port := 8081
	fmt.Printf("Account Service listening on port %d...\n", port)

	handleRequests(port)
}

func openDbConnection() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("db1.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database !!!")
	}
	return db
}

func createDbModels(db *gorm.DB) {
	err := db.AutoMigrate(models.Account{})
	if err != nil {
		panic("Failed to create database models !!!")
	}
}

func handleRequests(port int) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/account", handlers.CreateUser).Methods("POST")
	myRouter.HandleFunc("/account/balance", handlers.GetBalance).Methods("GET")
	myRouter.HandleFunc("/account/{id}/deposit", handlers.TransactionMoney).Methods("PATCH")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), myRouter))
}
