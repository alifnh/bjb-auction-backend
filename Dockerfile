FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/api

FROM alpine:3 AS production

WORKDIR /app

COPY --from=builder /app/main /app/
COPY .env.example /app/.env

EXPOSE 8000

ENTRYPOINT ["./main"]