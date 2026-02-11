package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	utils "tcc-test/api/utils"
)

// Validates the presence and integrity of a JWT token from the Authorization header.
func CheckBarrierHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract the Authorization header (expected format: "Bearer <token>")
		tokenString := c.GetHeader("Authorization")
		// Reject request if no token is provided
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing barrier token"})
			c.Abort()
			return
		}

		// Remove "Bearer " prefix to isolate the token
		tokenString = tokenString[len("Bearer "):]

		// Parse and validate the JWT token using the secret key
		claim, err := utils.Validate(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// Store token claims in Gin context for downstream access (e.g., user ID, roles)
		c.Set("claims", *claim)

		// Continue to next middleware or handler
		c.Next()
	}
}
