package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IPWhitelist(allowed []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		for _, a := range allowed {
			if ip == a {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusForbidden, gin.H{"error": "ip not allowed"})
		c.Abort()
	}
}