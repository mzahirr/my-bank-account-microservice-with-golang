package handlers

import (
	"account-service/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	// ::todo
	balance := 1000.0
	fmt.Fprintf(w, "Hesap bakiyesi: $%.2f", balance)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user models.Account

	if err := decoder.Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Yeni bir kullanıcı oluşturun.
	if newAccount, err := models.CreateAccount(user.FirstName, user.LastName, user.Balance); err != nil {
		if err != nil {
			fmt.Println("Hesap oluşturulamadı:", err)
			http.Error(w, "Hesap oluşturulamadı:", http.StatusInternalServerError)
			return
		} else {
			fmt.Println(newAccount)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	convertedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Bad Parameter:", err)
		http.Error(w, "Bad Parameter:", http.StatusBadRequest)
		return
	}

	db, err := gorm.Open(sqlite.Open("db1.sqlite"), &gorm.Config{})
	if err != nil {
		return
	}

	var user models.Account
	if err := db.First(&user, convertedId).Error; err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
