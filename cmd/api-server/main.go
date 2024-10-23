package main

import (
	apiserver "fastbin/internal/api-server"
	"fmt"
	"net/http"
)

func main() {
	port := 8080
	server := apiserver.NewAPIServer(port)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
