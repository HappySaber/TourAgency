# Используем официальный Go образ
FROM golang:1.23.2 AS builder

RUN mkdir /app

WORKDIR /app

# Устанавливаем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем весь код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/bin ./cmd/main.go

# Минимальный образ для запуска
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/bin /app/bin

COPY ./internal/migrations/ /docker-entrypoint-initdb.d/

# Копируем файл env из директории src
COPY .env .env 

# Указываем переменную окружения для порта
ENV PORT=8080

EXPOSE 8080

CMD ["/app/bin"]
