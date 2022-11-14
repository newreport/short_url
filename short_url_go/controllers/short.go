package controllers

import (
	"encoding/json"
	"short_url_go/models"

	"github.com/beego/beego/v2/core/logs"
)

// Operations about Shorts
type ShortController struct {
	BaseController
}

// @Title	create short
// @Summary	新增一个短链接
// @Description add one short url
// @Param	short		body 	models.Short		"body for short"
// @Success 200 {bool} add success
// @router / [post]
func (s *ShortController) CreateShort() {
	var short models.Short
	accInfo := s.analysisAccountClaims()
	json.Unmarshal(s.Ctx.Input.RequestBody, &short)
	short.FkUser = accInfo.ID
	s.Data["json"] = models.CreateShort(short, 6)
	s.ServeJSON()
}

// @Title	get shorts by page
// @Summary	分頁查詢
// @Date	2022-11-14
// @Auth	sfhj
// @Param	page	path	int	true	"頁"
// @Param	rows	path	int true	"行"
// @Param	sourceURL	query	string	false	"源url"
// @Param	targetURL	query	string false	"目标url（短链接）"
// @Param	isEnable	query	int	false	"是否启用"
// @Param	isExp	query	int	false	"是否过期"
// @Param	startAt	query	string	false	"创建时间start"
// @Param	endAt	query	string	false	"创建时间end"
// @Param	group	query	string flase	"分组名称"
// @Param	page	query	models.Page	"分页查询"
// @Success 200 bool
// @router /page/:page/rows/:rows [get]
func (s *ShortController) GetShortsByPage() {
	// accInfo := s.analysisAccountClaims()
	logs.Info("page: ", s.GetString(":page"))
	logs.Info("rows: ", s.GetString(":rows"))
	logs.Info("sourceURL: ", s.GetString("sourceURL"))
	s.Data["json"] = models.QueryPageShort()
	s.ServeJSON()
	// logs.Info(accInfo)
}
