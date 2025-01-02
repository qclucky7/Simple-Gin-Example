package middleware

import (
	locale "gin-quick-start/internal/components/locale"
	contextManager "gin-quick-start/internal/context"

	"github.com/gin-gonic/gin"
)

func Locale() gin.HandlerFunc {
	return func(c *gin.Context) {
		language := c.Request.Header.Get("Accept-Language")
		matchLocale := locale.LookupMatch(language)
		contextManager.SetLocale(c, matchLocale)
		c.Next()
	}
}
