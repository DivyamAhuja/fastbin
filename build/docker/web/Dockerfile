FROM golang:1.23-alpine AS builder

WORKDIR /app

RUN apk --no-cache add curl nodejs npm

RUN npm install -D tailwindcss

RUN go install github.com/a-h/templ/cmd/templ@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN templ generate


RUN npx tailwindcss -i ./internal/web/assets/css/input.css -o ./internal/web/assets/css/output.css

RUN go build -o web ./cmd/web

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/web /app/web

ENV GIN_MODE=release

EXPOSE 8080

ENTRYPOINT ["/app/web"]
