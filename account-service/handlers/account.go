package handlers

import (
	"account-service/models"
	"encoding/json"
	"fmt"
	"net/http"
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
	if err := models.CreateAccount(user.FirstName, user.LastName, user.Balance); err != nil {
		http.Error(w, "test", http.StatusInternalServerError)
		return
	}

	// Kullanıcıyı yanıtlayın.
	fmt.Fprintf(w, "Kullanıcı oluşturuldu: %+v", user)
}
