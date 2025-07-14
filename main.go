package main

import (
	routes "golangbackendauth/routes" // routes is a package name given to the files in the routes folder
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

	routes.AuthRoutes(router) // AuthRoutes is the main function in the routes package exporting the api endpoints from the router
	routes.UserRoutes(router) // UserRoutes is a function in the routes package

	// This line starts the Gin web server
	router.Run(":" + port)
}
