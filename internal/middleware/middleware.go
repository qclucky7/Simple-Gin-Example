package middleware

import (
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	r.Use(Cors())
	r.Use(RequestId())
	r.Use(Locale())
	r.Use(LogPrinter())
	r.Use(ApiErrorHandler())
}
