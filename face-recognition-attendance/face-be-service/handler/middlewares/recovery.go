package middlewares

import (
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"face-be-service/common/constants"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Error(string(debug.Stack()))
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": constants.SomethingWentWrong,
				})
			}
		}()

		c.Next()
	}
}
