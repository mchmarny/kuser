package main

import (
	"log"
	"os"
	"fmt"
	"context"

	"github.com/gin-gonic/gin"


)

const (
	defaultPort      = "8080"
	portVariableName = "PORT"
)

func main() {

	ctx := context.Background()
	initStore(ctx)

	// router
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// root & health
	r.GET("/", healthHandler)
	r.GET("/:id", getUserHandler)
	r.POST("/user", saveUserHandler)
	r.POST("/event", saveUserEventHandler)
	r.GET("/health", healthHandler)

	// port
	port := os.Getenv(portVariableName)
	if port == "" {
		port = defaultPort
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server starting: %s \n", addr)
	if err := r.Run(addr); err != nil {
		log.Fatal(err)
	}

}
