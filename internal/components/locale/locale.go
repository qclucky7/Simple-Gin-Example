package locale

import (
	"github.com/gookit/goutil/strutil"
	"golang.org/x/text/language"
)

var (
	DefaltLocale   = LanguageWrapper{language.Make("en-US"), "en-US"}
	SupportLocales = []LanguageWrapper{
		{language.Make("en-US"), "en-US"},
		{language.Make("zh-CN"), "zh-CN"},
	}
)

type LanguageWrapper struct {
	Tag    language.Tag
	Locale string
}

var localeMatcher language.Matcher

func SetUp() {
	tags := make([]language.Tag, len(SupportLocales))
	for _, data := range SupportLocales {
		tags = append(tags, data.Tag)
	}
	localeMatcher = language.NewMatcher(tags)
}

// 获取支持语言标签
func SupportLanageTags() []language.Tag {
	lanageTags := make([]language.Tag, 0)
	for _, item := range SupportLocales {
		lanageTags = append(lanageTags, item.Tag)
	}
	return lanageTags
}

// 查找匹配语言
func LookupMatch(locale string) LanguageWrapper {
	if strutil.IsEmpty(locale) {
		return DefaltLocale
	}
	tags, _, err := language.ParseAcceptLanguage(locale)
	if err != nil {
		return DefaltLocale
	}
	tag, _, _ := localeMatcher.Match(tags...)
	for _, tagWrapper := range SupportLocales {
		if tagWrapper.Tag == tag {
			return tagWrapper
		}
	}
	return DefaltLocale
}
