FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o api-server ./cmd/api-server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/api-server /app/api-server

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/app/api-server"]
