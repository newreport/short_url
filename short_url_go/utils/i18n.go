package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

func init() {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.MustLoadMessageFile("data/locale_en-US.toml")
	bundle.MustLoadMessageFile("data/locale_fr-FR.toml")
	bundle.MustLoadMessageFile("data/locale_ha-JP.toml")
	bundle.MustLoadMessageFile("data/locale_zh-CN.toml")

	loc = i18n.NewLocalizer(bundle, "locale_en-US")
}
