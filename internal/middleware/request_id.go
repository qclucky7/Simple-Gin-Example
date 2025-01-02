package middleware

import (
	contextManager "gin-quick-start/internal/context"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/goutil/strutil"
)

const (
	REQUEST_ID_KEY = "X-Request-Id"
)

func RequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(REQUEST_ID_KEY)
		if strutil.IsEmpty(requestId) {
			requestId = uuid.New().String()
		}
		contextManager.SetRequestId(c, requestId)
		c.Next()
		c.Writer.Header().Set(REQUEST_ID_KEY, requestId)
	}
}
