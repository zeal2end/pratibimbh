package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger returns a middleware that logs incoming HTTP requests
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		startTime := time.Now()

		// Get request path and method before calling next
		path := c.Request.URL.Path
		method := c.Request.Method

		// Get client IP
		clientIP := c.ClientIP()

		// Process request
		c.Next()

		// After request is processed
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		// Get response status
		status := c.Writer.Status()

		// Format and print log
		logMessage := fmt.Sprintf("[GIN] %v | %3d | %13v | %15s | %-7s %s",
			endTime.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			clientIP,
			method,
			path,
		)

		switch {
		case status >= 500:
			fmt.Printf("\033[31m%s\033[0m\n", logMessage) // Red for server errors
		case status >= 400:
			fmt.Printf("\033[33m%s\033[0m\n", logMessage) // Yellow for client errors
		case status >= 300:
			fmt.Printf("\033[36m%s\033[0m\n", logMessage) // Cyan for redirects
		default:
			fmt.Printf("\033[32m%s\033[0m\n", logMessage) // Green for success
		}

		// Log request body size if it's a POST/PUT request
		if method == "POST" || method == "PUT" {
			fmt.Printf("Request Body Size: %d bytes\n", c.Request.ContentLength)
		}
	}
}
