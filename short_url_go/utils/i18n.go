package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/beego/beego/logs"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

func init() {
	Bundle := i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	Bundle.LoadMessageFile("data/default.en.toml")
	Bundle.LoadMessageFile("data/default.zh.toml")
	logs.Info("开启测试")
	localizer := i18n.NewLocalizer(Bundle, "default.en")
	helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{TemplateData: "paramsError"})
	logs.Info(helloPerson)
	// Bundle.MustLoadMessageFile("data/locale_zh-CN.toml")
	// Bundle.MustLoadMessageFile("data/locale_fr-FR.toml")
	// Bundle.MustLoadMessageFile("data/locale_ha-JP.toml")
}
