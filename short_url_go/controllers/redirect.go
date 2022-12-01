package controllers

import (
	"short_url_go/utils"

	"github.com/beego/beego/v2/core/logs"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type RedirectController struct {
	BaseController
}

// func (r *RedirectController) Get() {
// 	url := r.Ctx.Input.Param(":url")
// 	logs.Info(url)
// 	r.Ctx.Redirect(302, "https://www.baidu.com/")

// }

func (r *RedirectController) Get() {
	// url := r.Ctx.Input.Param(":url")
	logs.Info("进入测试")
	localizer := i18n.NewLocalizer(utils.Bundle, "default.es")
	helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "paramsError"})
	logs.Info(helloPerson)
	r.Ctx.WriteString(helloPerson)

}
