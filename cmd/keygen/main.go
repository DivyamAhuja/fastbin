package main

import (
	keygen "fastbin/internal/keygen"
	"fmt"
	"log"
	"net"
)

func main() {
	port := 8081
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = keygen.NewKeygenServer().Serve(lis)
	if err != nil {
		log.Fatalf("grpc server error: %v", err)
	}
}
