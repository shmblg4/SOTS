FROM golang:1.20-alpine as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o tcp-server-service .

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /app/tcp-server-service .

CMD ["./tcp-server-service"]
