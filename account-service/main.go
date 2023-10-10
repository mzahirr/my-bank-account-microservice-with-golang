package main

import (
	"account-service/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/account/balance", handlers.GetBalance)
	http.HandleFunc("/account", handlers.CreateUser)

	port := 8081
	fmt.Printf("Account Service listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
