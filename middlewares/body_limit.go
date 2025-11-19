package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BodyLimit(maxBytes int64) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, maxBytes)
		if err := c.Request.ParseForm(); err != nil && !errors.Is(err, http.ErrNotMultipart) {
			c.JSON(http.StatusRequestEntityTooLarge, gin.H{"error": "body too large"})
			c.Abort()
			return
		}
		c.Next()
	}
}