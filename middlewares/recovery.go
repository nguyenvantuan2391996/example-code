package middlewares

import (
	"fmt"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))
				logrus.Error(fmt.Sprintf("detail panic %v", string(debug.Stack())))

				// alert

				ctx.AbortWithStatusJSON(500, gin.H{
					"error": "Something went wrong!",
				})
				return
			}
		}()

		ctx.Next()
	}
}
