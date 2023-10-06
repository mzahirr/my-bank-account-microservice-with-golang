// transaction-service/handlers/transaction.go

package handlers

import (
	"fmt"
	"net/http"
)

func TransferMoney(w http.ResponseWriter, r *http.Request) {
	// Para transferi işlemleri burada gerçekleştirilir
	// İlgili istek verilerini işleyin ve işlem sonucunu yanıt olarak döndürün
	fmt.Fprintf(w, "Para transferi başarıyla tamamlandı.")
}
