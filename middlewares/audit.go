package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuditLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		log.Printf("Audit: %s %s %s -> %d", c.Request.Method, c.Request.URL.Path, c.ClientIP(), c.Writer.Status())
	}
}