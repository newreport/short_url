package controllers

import (

)

// 302 後台跳轉
type RedirectController struct {
	BaseController
}

// func (r *RedirectController) Get() {
// 	logs.Info("进入测试")
// 	r.infos()
// 	domain := r.Ctx.Input.Host()
// 	user := models.QueryUserByDomain(domain)
// 	if user.ID > 0 {
// 		targetURL := r.Ctx.Input.Param(":shortURL")
// 		short := models.QueryShortByFKUserSourceURL(user.ID, targetURL)
// 		if len(short.ID) > 0 {
// 			r.Ctx.Redirect(302, short.SourceURL)
// 		}
// 	}
// 	// localizer := i18n.NewLocalizer(utils.Bundle, "default.es")
// 	// helloPerson := localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: "paramsError"})
// 	// logs.Info(helloPerson)
// 	// r.Ctx.WriteString(helloPerson)
// }
