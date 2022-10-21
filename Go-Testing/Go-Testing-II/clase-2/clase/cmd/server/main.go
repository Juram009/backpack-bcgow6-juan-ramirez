package main

import "github.com/gin-gonic/gin"

func main() {
	// Create router.
	router := gin.Default()

	// Add routes.
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Start and run the server.
	router.Run()
}
