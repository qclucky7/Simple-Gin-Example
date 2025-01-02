package contextManager

import (
	"gin-quick-start/internal/components/locale"
	Logger "gin-quick-start/internal/components/logger"
	models "gin-quick-start/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

const (
	REQUEST_ID = "x-request-id"
	SESSION    = "session"
	LOGGER     = "logger"
	AUTH_TOKEN = "token"
	LOCALE     = "locale"
)

func SetAuthToken(c *gin.Context, token string) {
	c.Set(AUTH_TOKEN, token)
}

func GetAuthToken(c *gin.Context) string {
	return c.GetString(AUTH_TOKEN)
}

func SetRequestId(c *gin.Context, requestId string) {
	c.Set(REQUEST_ID, requestId)
}

func GetRequestId(c *gin.Context) string {
	return c.GetString(REQUEST_ID)
}

func SetLocale(c *gin.Context, locale locale.LanguageWrapper) {
	c.Set(LOCALE, locale)
}

func GetLocale(c *gin.Context) locale.LanguageWrapper {
	_locale, exist := c.Get(LOCALE)
	if !exist {
		return locale.DefaltLocale
	}
	return _locale.((locale.LanguageWrapper))
}

func SetSession(c *gin.Context, session *models.Session) {
	c.Set(SESSION, session)
}

func GetSession(c *gin.Context) *models.Session {
	session, exist := c.Get(SESSION)
	if exist {
		return session.(*models.Session)
	}
	return &models.Session{}
}

func SetLogger(c *gin.Context, entry *logrus.Entry) {
	c.Set(LOGGER, entry)
}

func GetLogger(c *gin.Context) *logrus.Entry {
	if c == nil {
		return Logger.Logger.WithFields(logrus.Fields{})
	}
	instance, exist := c.Get(LOGGER)
	if exist {
		return instance.(*logrus.Entry)
	}
	return Logger.Logger.WithFields(logrus.Fields{})
}
