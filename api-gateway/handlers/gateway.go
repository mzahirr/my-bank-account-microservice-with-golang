// api-gateway/handlers/gateway.go

package handlers

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyToAccountService(w http.ResponseWriter, r *http.Request) {
	// Account Service URL
	accountServiceURL := "http://account-service:8081"

	// İstek yönlendirme
	proxy := httputil.NewSingleHostReverseProxy(parseURL(accountServiceURL))
	proxy.ServeHTTP(w, r)
}

func ProxyToTransactionService(w http.ResponseWriter, r *http.Request) {
	// Transaction Service URL
	transactionServiceURL := "http://transaction-service:8082"

	// İstek yönlendirme
	proxy := httputil.NewSingleHostReverseProxy(parseURL(transactionServiceURL))
	proxy.ServeHTTP(w, r)
}

func parseURL(rawURL string) *url.URL {
	u, err := url.Parse(rawURL)
	if err != nil {
		fmt.Printf("URL parsing error: %v\n", err)
	}
	return u
}
