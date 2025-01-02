package middleware

import (
	"fmt"
	errors "gin-quick-start/internal/errors"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

var (
	INTERNAL_REQUEST_HEADER_KEY = "X-Internal-Request"
)

func RateLimit(requestsPerSecond float64) gin.HandlerFunc {
	limiter := tollbooth.NewLimiter(requestsPerSecond, nil)
	return func(c *gin.Context) {
		requestMatchPath := c.FullPath()
		requestMethod := c.Request.Method
		err := tollbooth.LimitByKeys(limiter, []string{requestMethod, requestMatchPath})
		c.Writer.Header().Add("X-Rate-Limit-Limit", fmt.Sprintf("%.2f", limiter.GetMax()))
		c.Writer.Header().Add("X-Rate-Limit-Duration", "1")
		c.Writer.Header().Add("X-Rate-Limit-Request-Forwarded-For", c.GetHeader("X-Forwarded-For"))
		if err != nil {
			panic(errors.RATE_LIMIT_ERROR)
		}
		c.Next()
	}
}
