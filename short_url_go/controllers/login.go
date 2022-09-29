package controllers

import beego "github.com/beego/beego/v2/server/web"

// Operations about Users
type LoginController struct {
	beego.Controller
}

//https://cloud.tencent.com/developer/article/1557075

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (this *LoginController) Logout() {
	this.Data["json"] = "logout success"
	this.ServeJSON()
}

// @Title logout1
// @Description Logs1 out current logged in user session
// @Success 200 {string} logout1 success
// @router /logout1 [get]
func (this *LoginController) Logout1() {
	this.Data["json"] = "logout success"
	this.ServeJSON()
}
