package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/spf13/viper"

	"face-be-service/common/constants"
)

func APIKeyAuthentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set(constants.RequestIDField, uuid.NewString())
		if viper.GetString(constants.APIKey) != ctx.Request.Header.Get(constants.XAPIKeyHeader) {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "You are not authorized to perform the action",
			})
			return
		}

		ctx.Next()
	}
}
