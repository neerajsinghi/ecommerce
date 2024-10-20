package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger is a middleware function that logs the details of each request
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Process request
		c.Next()

		// Calculate resolution time
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get status code
		statusCode := c.Writer.Status()

		// Log details
		log.Printf("Status: %d | Latency: %v | Method: %s | Path: %s",
			statusCode, latency, c.Request.Method, c.Request.URL.Path)
	}
}
