package main

import (
	keygen "fastbin/internal/keygen"
	"fastbin/internal/pkg/env"
	"fmt"
	"log"
	"net"
	"strconv"
)

func main() {
	port, err := strconv.Atoi(env.GetEnv("KEYGEN_INTERNAL_PORT", "8080"))
	if err != nil {
		log.Fatalf("error listening port: %v, err: %v", port, err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	err = keygen.NewKeygenServer().Serve(lis)
	if err != nil {
		log.Fatalf("grpc server error: %v", err)
	}
}
