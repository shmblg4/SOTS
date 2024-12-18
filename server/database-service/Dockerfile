FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o db-service .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/db-service .

CMD ["./db-service"]
