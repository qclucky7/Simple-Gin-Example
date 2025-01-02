package middleware

import (
	"fmt"
	Logger "gin-quick-start/internal/components/logger"
	contextManager "gin-quick-start/internal/context"
	errors "gin-quick-start/internal/errors"
	models "gin-quick-start/internal/models"
	"net/http"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/arrutil"
)

func handelApiError(c *gin.Context, error errors.APIError) models.Result[string] {
	i18nMessage, err := ginI18n.GetMessage(c, error.Message)
	if err != nil {
		i18nMessage = error.Message
	}
	return models.Result[string]{
		Status: error.Code,
		Data:   fmt.Sprintf(i18nMessage, arrutil.StringsToAnys(error.Extra)...),
	}
}

func ApiErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger := contextManager.GetLogger(c)
				if logger != nil {
					logger.Error(err)
				} else {
					Logger.Logger.Error(err)
				}
				switch e := err.(type) {
				case errors.APIError:
					c.AbortWithStatusJSON(
						http.StatusOK,
						handelApiError(c, e),
					)
					return
				default:
					contextManager.GetLogger(c).Error(err)
					c.AbortWithStatusJSON(
						http.StatusOK,
						handelApiError(c, errors.SERVER_ERROR),
					)
					return
				}
			}
		}()
		c.Next()
		for _, err := range c.Errors {
			switch e := err.Err.(type) {
			case errors.APIError:
				c.AbortWithStatusJSON(
					http.StatusOK,
					handelApiError(c, e),
				)
				return
			default:
				contextManager.GetLogger(c).Error(err)
				c.AbortWithStatusJSON(
					http.StatusOK,
					handelApiError(c, errors.SERVER_ERROR),
				)
				return
			}
		}
	}
}
