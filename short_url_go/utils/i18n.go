package utils

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

var Bundle *i18n.Bundle

func init() {
	Bundle = i18n.NewBundle(language.English)
	Bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	Bundle.LoadMessageFile("../data/go/data/en.toml")
	Bundle.LoadMessageFile("../data/go/data/zh.toml")
	Bundle.MustParseMessageFileBytes([]byte(`
	HelloWorld = "Hola Mundo!"
	`), "es.toml")
	// logs.Info("开启测试")
	// str, _ := os.Getwd()
	// logs.Info(str)
	// localizer := i18n.NewLocalizer(Bundle, "en")
	// logs.Info(localizer)
	// logs.Info("开启i18n")
	// helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "HelloWorld"})
	// logs.Info(helloPerson)
	// Bundle.MustLoadMessageFile("data/locale_zh-CN.toml")
	// Bundle.MustLoadMessageFile("data/locale_fr-FR.toml")
	// Bundle.MustLoadMessageFile("data/locale_ha-JP.toml")
}
