package handlers

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

// ProxyHandler creates a Gin handler function that forwards requests to a target URL.
// It directly uses the provided targetURL for the request, which is expected to be
// the exact URL for the downstream service.
func ProxyHandler(targetURL string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Parse the target URL provided to ProxyHandler.
		// This targetURL already contains the correct path for the downstream service.
		parsedTargetURL, err := url.Parse(targetURL)
		if err != nil {
			log.Printf("Error parsing target URL %s: %v", targetURL, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Create a new HTTP request using the parsedTargetURL directly.
		// We no longer use c.Request.URL here, as parsedTargetURL already has the correct path
		// for the downstream service after processing by DynamicProxyHandler.
		req, err := http.NewRequest(c.Request.Method, parsedTargetURL.String(), c.Request.Body)
		if err != nil {
			log.Printf("Error creating proxy request: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}

		// Copy headers from the original request to the new request.
		// This is crucial for forwarding things like Authorization headers.
		for name, values := range c.Request.Header {
			for _, value := range values {
				req.Header.Add(name, value)
			}
		}

		// If the AuthMiddleware has set a user ID in the context,
		// add it as a custom header to the proxied request.
		// This allows downstream services to trust the API Gateway's authentication.
		if userID, exists := c.Get("userID"); exists {
			if idStr, ok := userID.(string); ok {
				req.Header.Set("X-User-ID", idStr)
			}
		}

		// Log the final URL that will be sent to the downstream service for debugging.
		log.Print(req.URL.String(), " - Final proxied request URL")

		// Execute the request using a default HTTP client.
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Printf("Error proxying request to %s: %v", parsedTargetURL.String(), err)
			c.JSON(http.StatusBadGateway, gin.H{"error": "Cannot reach upstream service"})
			return
		}
		defer resp.Body.Close()

		// Copy headers from the proxy response to the original response writer.
		for name, values := range resp.Header {
			for _, value := range values {
				c.Writer.Header().Add(name, value)
			}
		}

		// Set the status code from the proxy response.
		c.Status(resp.StatusCode)

		// Copy the response body from the proxy response to the original response writer.
		if _, err := io.Copy(c.Writer, resp.Body); err != nil {
			log.Printf("Error copying proxy response body: %v", err)
			// Note: We can't change the status code here if headers have already been written.
		}
	}
}

// DynamicProxyHandler creates a Gin handler that proxies requests to a target service.
// It takes the base URL of the target service and the path prefix used by the API Gateway.
// It strips the gateway prefix from the incoming request path before forwarding.
func DynamicProxyHandler(serviceBaseURL, gatewayPathPrefix string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the full path from the incoming request (e.g., "/users/login" or "/todo/todo/123")
		incomingPath := c.Request.URL.Path

		// Strip the gatewayPathPrefix to get the path expected by the downstream service.
		// Example: if incomingPath is "/todo/todo/123" and gatewayPathPrefix is "/todo",
		// then servicePath will be "/todo/123".
		servicePath := strings.TrimPrefix(incomingPath, gatewayPathPrefix)

		// Ensure the servicePath starts with a '/' if it's not empty.
		// This is important for correctly forming the target URL.
		if servicePath != "" && !strings.HasPrefix(servicePath, "/") {
			servicePath = "/" + servicePath
		}

		// Construct the full target URL for the downstream service.
		// This URL will be passed to ProxyHandler.
		targetURL := serviceBaseURL + servicePath + "?" + c.Request.URL.RawQuery

		// Log the incoming path and the target URL constructed by DynamicProxyHandler for debugging.
		log.Print(incomingPath, " - incoming path, ", targetURL, " - in dynamic function")

		// Now, use the existing ProxyHandler with the dynamically constructed targetURL.
		ProxyHandler(targetURL)(c)
	}
}

