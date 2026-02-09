package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Import and register auth router
	// authRouter := adapters.NewAuthRouter()
	// authRouter.RegisterRoutes(r)

	// Health check endpoint
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "healthy"})
	})

	r.Run(":8080")
}
