package middleware

import (
	contextManager "gin-quick-start/internal/context"
	errors "gin-quick-start/internal/errors"
	models "gin-quick-start/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
)

const (
	AUTHORIZATION = "Authorization"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader(AUTHORIZATION)
		if strutil.IsEmpty(token) {
			panic(errors.AUTH_TOKEN_UNAUTHORIZED_ACCESS)
		}
		contextManager.SetAuthToken(c, token)
		contextManager.SetSession(c, &models.Session{
			Account:   "",
			AccountId: "",
			Powers:    []string{},
		})
		c.Next()
	}
}
