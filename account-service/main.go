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

	db, err := gorm.Open(sqlite.Open("db1.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Veritabanına bağlanılamadı!")
	}

	// Account modeli için tabloyu oluştur veya güncelle
	err2 := db.AutoMigrate(models.Account{})
	if err2 != nil {
		panic(err2)
	}

	port := 8081
	fmt.Printf("Account Service listening on port %d...\n", port)

	handleRequests(port)
}

func handleRequests(port int) {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/account", handlers.CreateUser).Methods("POST")
	myRouter.HandleFunc("/account/balance", handlers.GetBalance).Methods("GET")
	myRouter.HandleFunc("/account/{id}/deposit", handlers.TransactionMoney).Methods("PATCH")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), myRouter))
}
