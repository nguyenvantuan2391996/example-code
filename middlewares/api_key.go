package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func APIKeyMiddleware(validKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-Key")
		if key != validKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid api key"})
			c.Abort()
			return
		}
		c.Next()
	}
}