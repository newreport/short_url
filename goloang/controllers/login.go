package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	beego.Controller
}

// @Title GetAll
// @Success 200 string
// @router /getall [get]
func (this *LoginController) GetAll() {
	this.Data["json"] = "i'm controller"
	this.ServeJSON()
}

func (this *LoginController) IsInit() {

}
