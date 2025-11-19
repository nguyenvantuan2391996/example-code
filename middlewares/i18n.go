package main

import "github.com/gin-gonic/gin"

func I18N(defaultLang string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("Accept-Language")
		if lang == "" {
			lang = defaultLang
		}
		c.Set("lang", lang)
		c.Next()
	}
}