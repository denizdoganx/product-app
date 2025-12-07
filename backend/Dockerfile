# -----------------------------
# 1. Build aşaması
# -----------------------------
FROM golang:1.24-alpine AS builder

# Gerekli paketler
RUN apk add --no-cache git

WORKDIR /app

# Go mod dosyaları kopyalanır
COPY go.mod go.sum ./

# Modüller çekilir
RUN go mod download

# Tüm proje kopyalanır
COPY . .

# Go binary build edilir
RUN go build -o server ./cmd/main.go

# -----------------------------
# 2. Run aşaması (çok küçük image)
# -----------------------------
FROM alpine:3.20

# Sertifika ekliyoruz (https çağrıları için gerekli)
RUN apk add --no-cache ca-certificates

WORKDIR /app

# Builder'dan binary kopyalanır
COPY --from=builder /app/server .

# Port
EXPOSE 8080

# Başlat
CMD ["./server"]
