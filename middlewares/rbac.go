package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RBAC(requiredRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, ok := c.Get("jwt_claims")
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token claims"})
			c.Abort()
			return
		}


		mapClaims := claims.(jwt.MapClaims)


		role, ok := mapClaims["role"].(string)
		if !ok {
			c.JSON(http.StatusForbidden, gin.H{"error": "missing role"})
			c.Abort()
			return
		}


		for _, r := range requiredRoles {
			if r == role {
				c.Next()
				return
			}
		}


		c.JSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
		c.Abort()
	}
}