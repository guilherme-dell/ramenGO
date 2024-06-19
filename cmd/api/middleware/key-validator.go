package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiKeyAuth(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("x-api-key")

		if apiKey != validKey {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "x-api-key header missing"})
			return
		}

		c.Next()
	}
}
