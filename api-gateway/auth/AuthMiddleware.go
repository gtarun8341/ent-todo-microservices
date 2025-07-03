package middleware

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == ""{
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}
		token := parts[1]
		c.Set("sessionId",token)

		resp, err := http.Get(fmt.Sprintf("http://localhost:8001/session?token=%s", token))
        if err != nil || resp.StatusCode != http.StatusOK {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
            c.Abort()
            return
        }
        defer resp.Body.Close()
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
			return
		}
		log.Println("Response body:", string(bodyBytes))

		var result struct {
			UserId string `json:"userId"`
		}
		if err := json.Unmarshal(bodyBytes, &result); err != nil || result.UserId == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid session"})
			c.Abort()
			return
		}

		c.Set("userId", result.UserId)
		c.Next()
	}

}