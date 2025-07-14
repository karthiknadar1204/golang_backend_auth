package routes

import (
	controller "golangbackendauth/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controller.Signup()) //-> This is the endpoint for the signup route
	incomingRoutes.POST("users/login", controller.Login()) //-> This is the endpoint for the login route
}

// incomingRoutes *gin.Engine is key to understanding how Gin works in Go.
// What is *gin.Engine?
// 🔹 gin.Engine is:
// The main object in Gin that powers your whole server.

// It’s like:
// The traffic controller for your API

// The master router
// The thing that listens to requests and sends responses

// 🔹 *gin.Engine (with the asterisk) means:
// This is a pointer to a gin.Engine object.

// Why a pointer? Because:
// It allows you to modify the engine inside a function.
// It's more efficient (Go passes a reference, not a full copy).


