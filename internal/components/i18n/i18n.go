package i18n

import (
	"encoding/json"
	"gin-quick-start/internal/components/locale"

	ginI18n "github.com/gin-contrib/i18n"
	"github.com/gin-gonic/gin"
)

func SetUp(r *gin.Engine) {
	r.Use(ginI18n.Localize(ginI18n.WithBundle(&ginI18n.BundleCfg{
		DefaultLanguage:  locale.DefaltLocale.Tag,
		FormatBundleFile: "json",
		AcceptLanguage:   locale.SupportLanageTags(),
		RootPath:         "./internal/components/i18n/message/",
		UnmarshalFunc:    json.Unmarshal,
	})))
}
