FROM golang:1.23.1-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/main.go --output docs

RUN go build -o notification-service cmd/main.go

# RUNTIME IMAGE
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/notification-service .
COPY .env .

EXPOSE 8080

CMD ["./notification-service"]