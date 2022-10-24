package controllers

import (
	"encoding/json"
	"short_url_go/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Shorts
type ShortController struct {
	beego.Controller
}

// @Title create short
// @Summary 新增一个短链接
// @Description add one short url
// @Param	short		body 	models.Short		"body for short"
// @Success 200 {bool} add success
// @Failure 403 not role
// @router / [post]
func (s *ShortController) CreateShort() {
	var short models.Short
	json.Unmarshal(s.Ctx.Input.RequestBody, &short)
	s.Data["json"] = models.CreateShort(short, 6)
	s.ServeJSON()
}
