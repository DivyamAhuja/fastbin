package apiserver

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

func hello_world(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World",
	})
}
