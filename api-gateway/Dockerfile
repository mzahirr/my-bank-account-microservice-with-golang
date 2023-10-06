# Resmi Go görüntüsünü temel alın
FROM golang:1.17

# Çalışma dizinini belirle
WORKDIR /app

# Kaynak kodları kopyala
COPY . .

# Uygulamayı derle
RUN go build -o api-gateway

# Uygulamayı çalıştır
CMD ["./api-gateway"]
