package apiserver

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	pb "fastbin/api/keygen"

	"fastbin/internal/pkg/env"
	paste "fastbin/internal/pkg/paste"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Paste = paste.Paste

type APIServer struct {
	db *gorm.DB
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	dbport   = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func NewAPIServer(port int) *http.Server {
	dbString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, database, dbport)
	db, err := gorm.Open(postgres.Open(dbString), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}

	db.AutoMigrate(&Paste{})

	s := APIServer{db: db}

	r := gin.Default()
	r.POST("/write", s.write)
	r.GET("/read/:key", s.read)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: r,
	}

	return server
}

func (as *APIServer) read(gc *gin.Context) {
	key, _ := gc.Params.Get("key")

	var paste Paste
	res := as.db.First(&paste, "id = ?", key)
	if res.Error != nil {
		gc.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found.",
			"text":  "",
		})
	} else {
		gc.JSON(http.StatusOK, gin.H{
			"text": paste.Text,
		})
	}
}

type WriteRequestBody struct {
	Text string
}

func (as *APIServer) write(gc *gin.Context) {
	var requestBody WriteRequestBody
	if err := gc.BindJSON(&requestBody); err != nil {
		fmt.Println(err)
		gc.JSON(http.StatusBadRequest, gin.H{
			"key": "",
		})
	}

	key, err := as.try_write(requestBody.Text)
	for tries := 0; err != nil && tries < 5; tries++ {
		key, err = as.try_write(requestBody.Text)
	}

	if err != nil {
		gc.JSON(http.StatusInternalServerError, gin.H{
			"key": "",
		})
	} else {
		gc.JSON(http.StatusOK, gin.H{
			"key": key,
		})
	}
}

func (as *APIServer) try_write(data string) (string, error) {
	grpcServerURL := env.GetEnv("KEYGEN_HOST", "localhost") + ":" + env.GetEnv("KEYGEN_PORT", "8080")
	conn, err := grpc.NewClient(grpcServerURL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return "", err
	}
	defer conn.Close()

	c := pb.NewKeygenClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	r, err := c.GenerateKey(ctx, &pb.Empty{})
	cancel()
	if err != nil {
		return "", err
	}

	key := r.Value
	res := as.db.Create(&Paste{ID: key, Text: data})
	if res.Error != nil {
		return "", res.Error
	}
	return key, nil
}
