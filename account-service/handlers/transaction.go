package handlers

import (
	"account-service/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func TransactionMoney(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	id := vars["id"]

	convertedId, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Bad Parameter:", err)
		http.Error(w, "Bad Parameter:", http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var transaction models.Transaction

	if err := decoder.Decode(&transaction); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if updatedAccount, err := models.TransactionMoney(uint(convertedId), transaction); err != nil {
		if err != nil {
			fmt.Println("İşlem yapılamadı:", err)
			http.Error(w, "\"İşlem yapılamadı:", http.StatusInternalServerError)
			return
		} else {
			fmt.Println(updatedAccount)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(transaction)
}
