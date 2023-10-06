package handlers

import (
	"fmt"
	"net/http"
)

func GetBalance(w http.ResponseWriter, r *http.Request) {
	// ::todo
	balance := 1000.0
	fmt.Fprintf(w, "Hesap bakiyesi: $%.2f", balance)
}
