package routers

import (
	api "gin-quick-start/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	api.SetupRoutes(r)
}
