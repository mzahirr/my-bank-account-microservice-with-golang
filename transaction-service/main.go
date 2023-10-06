package main

import (
	"fmt"
	"log"
	"net/http"
	"transaction-service/handlers"
)

func main() {
	http.HandleFunc("/transaction/transfer", handlers.TransferMoney)

	port := 8082
	fmt.Printf("Transaction Service listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
