package main

import (
	"fmt"
	cmd "gin-quick-start/cmd"
	routers "gin-quick-start/internal/api"
	"gin-quick-start/internal/components"
	"gin-quick-start/internal/configs"
	"gin-quick-start/internal/middleware"
	"strconv"

	_ "gin-quick-start/docs"
	swagger "gin-quick-start/internal/components/swagger"

	"github.com/gin-gonic/gin"
)

// @title			Web Server Documentation
// @version			1.0.0
// @host 			127.0.0.1:9000
// @schemes			http

// @tag.name		API
// @tag.description Web API

func main() {
	cmd.Execute()
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	components.SetUpComponents(engine)
	middleware.Setup(engine)
	routers.Setup(engine)
	swagger.SetUp(engine)
	err := engine.Run(":" + strconv.Itoa(configs.GetConfiguration().Port))
	if err != nil {
		panic(err)
	}
	fmt.Printf("Server Start on %d", configs.GetConfiguration().Port)
}
