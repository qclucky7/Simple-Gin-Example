package handlers

import (
	"gin-quick-start/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	//login
	r.POST("/v1/login", UserLogin)

	//api
	api := r.Group("/v1").
		Use(middleware.TokenAuth())
	{
		api.GET("/hello-world", HelloWord)
	}
}
