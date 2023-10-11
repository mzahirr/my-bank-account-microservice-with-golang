package main

import (
	"account-service/handlers"
	"account-service/models"
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/account/balance", handlers.GetBalance)
	http.HandleFunc("/account", handlers.CreateUser)

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
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
