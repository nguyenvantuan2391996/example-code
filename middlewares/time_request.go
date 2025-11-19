package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func TimeRequestDetector(threshold time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		elapsed := time.Since(start)
		if elapsed > threshold {
			log.Printf("Time Request [%s] %s %s: %s", c.Request.Method, c.Request.URL.Path, c.ClientIP(), elapsed)
		}
	}
}