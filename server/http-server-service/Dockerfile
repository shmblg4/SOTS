# Этап сборки
FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o http-server-service .

# Финальный этап
FROM ubuntu:22.04

WORKDIR /root/
# Копируем исполняемый файл
COPY --from=builder /app/http-server-service .
# Копируем статические файлы
COPY --from=builder /app/static /root/static

# Устанавливаем права на выполнение
RUN chmod +x http-server-service

# Запуск сервера
CMD ["./http-server-service"]
