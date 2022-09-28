package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type ShortResult struct {
	Id      int64
	
	UrlSort string
}
type ShortController struct {
	beego.Controller
}

// @Title GetList
// @Success 200 string
// @router /getlist [get]
func (this *ShortController) GetList() {
	this.Data["json"] = "i'm controller"
	this.ServeJSON()
}

// @Title AddOne
// @Success 200 string
// @router /:url [post]
func (this *ShortController) AddOne() {
	oldUrl := this.GetString(":url")
	if oldUrl != "" {

	}
	this.Data["json"] = true
	this.ServeJSON()
}

// @Title AddOne
// @Success 200 bool
// @router /editone [put]
func (this *ShortController) EditOne() {

}

func (this *ShortController) DeleteOne() {

}
