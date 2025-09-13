# Используем официальный образ Golang
FROM golang:1.22.5 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и устанавливаем зависимости (это ускоряет сборку)
COPY go.mod ./
RUN go mod download

# Копируем все исходники в контейнер
COPY . .

# Собираем бинарный файл
RUN go build -o inDriveHack ./cmd/inDriveHack/main.go

# Используем минимальный образ для финального контейнера
FROM debian:bookworm-slim

WORKDIR /app

# Копируем бинарник из builder-контейнера
COPY --from=builder /app/inDriveHackBackend /app/inDriveHackBackend

# Открываем порт
EXPOSE 8080

# Запускаем приложение
CMD ["/app/inDriveHackBackend"]