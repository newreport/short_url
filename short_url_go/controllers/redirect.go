﻿package controllers

import "github.com/beego/beego/v2/core/logs"

type RedirectController struct {
	BaseController
}

func (r *RedirectController) Get() {
	url := r.Ctx.Input.Param(":url")
	logs.Info(url)
	r.Ctx.Redirect(302, "https://www.baidu.com/")
	
	t := gi18n.New()
}

