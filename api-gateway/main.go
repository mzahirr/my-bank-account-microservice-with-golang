// api-gateway/main.go
package main

import (
	"api-gateway/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/account", handlers.ProxyToAccountService)
	http.HandleFunc("/transaction/transfer", handlers.ProxyToTransactionService)

	port := 8080
	fmt.Printf("API Gateway listening on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
