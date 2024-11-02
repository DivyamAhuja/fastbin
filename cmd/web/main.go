package main

import (
	"fastbin/internal/pkg/env"
	webserver "fastbin/internal/web"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(env.GetEnv("WEB_INTERNAL_PORT", "8080"))
	if err != nil {
		log.Fatalf("error listening port: %v, err: %v", port, err)
	}
	server := webserver.NewServer(port)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
