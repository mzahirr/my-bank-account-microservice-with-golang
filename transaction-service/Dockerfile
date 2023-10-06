# Resmi Go görüntüsünü temel alın
FROM golang:1.17

# Çalışma dizinini belirle
WORKDIR /app

# Kaynak kodları kopyala
COPY . .

# Uygulamayı derle
RUN go build -o transaction-service

# Uygulamayı çalıştır
CMD ["./transaction-service"]
