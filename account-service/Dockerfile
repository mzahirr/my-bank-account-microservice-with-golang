# Resmi Go görüntüsünü temel alın
FROM golang:1.21

# Çalışma dizinini belirle
WORKDIR /app

# Kaynak kodları ve go mod dosyasını kopyala
COPY go.mod go.sum ./
RUN go mod download cmd/account-service/main.go

# Kaynak kodları kopyala
COPY . .

RUN go build -o account-service

RUN chmod +x account-service
CMD ["./account-service"]
