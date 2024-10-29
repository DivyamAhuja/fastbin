FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o keygen ./cmd/keygen

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/keygen /app/keygen

EXPOSE 8080

ENTRYPOINT ["/app/keygen"]
