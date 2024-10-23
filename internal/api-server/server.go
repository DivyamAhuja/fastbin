package apiserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	pb "fastbin/api/keygen"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewAPIServer(port int) *http.Server {
	r := gin.Default()
	r.GET("/", hello_world)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	return server
}

func hello_world(gc *gin.Context) {
	grpcServerURL := "localhost:8081"
	conn, err := grpc.NewClient(grpcServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewKeygenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GenerateKey(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	gc.JSON(http.StatusOK, gin.H{
		"message": "Hello, World",
		"key":     r.Value,
	})

}
