package handlers

import (
	"gin-quick-start/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/v1").
		Use(middleware.TokenAuth())
	{
		api.GET("/hello-world", HelloWord)
	}
}
