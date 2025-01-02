package components

import (
	http "gin-quick-start/internal/components/http"
	i18n "gin-quick-start/internal/components/i18n"
	locale "gin-quick-start/internal/components/locale"
	logger "gin-quick-start/internal/components/logger"

	"github.com/gin-gonic/gin"
)

func SetUpComponents(c *gin.Engine) {
	logger.SetUp()
	http.SetUp()
	locale.SetUp()
	i18n.SetUp(c)
}
