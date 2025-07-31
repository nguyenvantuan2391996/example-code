package middlewares

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func Recovery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(string(debug.Stack()))
				logger.GetInstance().GetLogger().Error(fmt.Sprintf("detail panic %v", string(debug.Stack())))

				// alert

				responseAPI := vbd_response.NewResponse(ctx)
				ctx.AbortWithStatusJSON(500, gin.H{
    				"error": "Something went wrong!",
				})
				return
			}
		}()

		ctx.Next()
	}
}