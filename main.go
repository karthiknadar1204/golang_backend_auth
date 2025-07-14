package main

import (
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT") //-> This is how you read environment variables in Go
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
}
