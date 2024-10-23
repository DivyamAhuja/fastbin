package main

import (
	apiserver "fastbin/internal/api-server"
	"fmt"
	"net/http"
)

func main() {
	server := apiserver.NewAPIServer(8080)
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}
}
