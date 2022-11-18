package controllers

import (
	"encoding/json"
	"short_url_go/models"
)

// Operations about Shorts
type ShortController struct {
	BaseController
}

// @Title	create short
// @Summary	新增一个短链接
// @Description add one short url
// @Param	short		body 	models.Short		"body for short"
// @Success 200	{string}	add success
// @Success 401	{string} 	Insufficient user permissions
// @router / [post]
func (s *ShortController) CreateShort() {
	var short models.Short
	accInfo := s.analysisAccountClaims()
	json.Unmarshal(s.Ctx.Input.RequestBody, &short)
	short.FKUser = accInfo.ID
	if models.CreateShort(short, 6) {
		s.Ctx.WriteString("创建成功")
	} else {
		s.Ctx.WriteString("创建失败")
	}
}

// @Title	get shorts by page
// @Summary	分頁查詢
// @Date	2022-11-14
// @Auth	sfhj
// @Param	page	query	models.Page	true	分页
// @Param	query	query	models.ShortPageQuery	false	查询参数
// @Success 200
// @router / [get]
func (s *ShortController) GetShortsByPage() {
	accInfo := s.analysisAccountClaims()
	var err error
	var page models.Page
	page.Offset, err = s.GetInt("offset")
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString("请求参数类型错误")
	}
	page.Lmit, err = s.GetInt("limit")
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString("请求参数类型错误")
	}
	page.Sort = analysisOrderBy(s.GetString("sort"))
	var query models.ShortPageQuery
	query.SourceURL = s.GetString("source_url")
	query.TargetURL = s.GetString("target_url")
	query.FKUser = accInfo.ID
	query.ShortGroup = s.GetString("group")
	query.IsEnable = s.GetString("is_enable")
	query.ExpireAt = s.GetString("exp")
	query.CreatedAt = s.GetString("crt")
	query.UpdatedAt = s.GetString("upt")
	query.DeletedAt = s.GetString("del")
	result, count, err := models.QueryPageShorts(query, page)
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString("请求参数错误")
	} else {
		s.Data["json"] = map[string]interface{}{
			"count": count,
			"data":  result,
		}
		s.ServeJSON()
	}
}
