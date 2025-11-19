package middlewares

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var rateMap = sync.Map{}

type rateInfo struct {
	Count int
	TS    time.Time
}

func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		v, _ := rateMap.LoadOrStore(ip, &rateInfo{Count: 0, TS: now})
		info := v.(*rateInfo)

		if now.Sub(info.TS) > window {
			info.Count = 0
			info.TS = now
		}

		info.Count++
		if info.Count > limit {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "rate limit exceeded"})
			c.Abort()
			return
		}

		c.Next()
	}
}
