package components

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	NAME = "Web Server Documentation"
)

func SetUp(c *gin.Engine) {
	c.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler,
		func(c *ginSwagger.Config) {
			c.Title = NAME
		},
	))
}
