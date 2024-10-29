package main

import (
	apiserver "fastbin/internal/api-server"
	"fastbin/internal/pkg/env"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(env.GetEnv("API_INTERNAL_PORT", "8080"))
	if err != nil {
		log.Fatalf("error listening port: %v, err: %v", port, err)
	}
	server := apiserver.NewAPIServer(port)
	err = server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
