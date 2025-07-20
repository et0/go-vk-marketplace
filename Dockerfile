# Шаг сборки
FROM golang:1.24 AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o go-vk-marketplace cmd/api/main.go

# Финальный образ
FROM alpine:latest
WORKDIR /app

# Установка зависимостей для работы с PostgreSQL
RUN apk add --no-cache libc6-compat

COPY --from=builder /app/go-vk-marketplace .
COPY --from=builder /app/config/local.yaml ./config/local.yaml

EXPOSE 8080
CMD ["./go-vk-marketplace"]