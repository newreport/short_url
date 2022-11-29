package controllers

import (
	"encoding/json"
	"short_url_go/models"
	"strconv"
)

// Operations about Shorts
type ShortController struct {
	BaseController
}

// @Title	create short
// @Summary	新增一个短链接
// @Description add one short url
// @Param	len	path	len	"默认长度"
// @Param	body	body	models.AddEditShort	true	"链接"
// @Success 200	{string}	"add success"
// @Failure 200	{string} 	"add fail"
// @router / [post]
func (s *ShortController) CreateShort() {
	accInfo := s.analysisAccountClaims()
	var dtoShort models.AddEditShort
	json.Unmarshal(s.Ctx.Input.RequestBody, &dtoShort)
	var short models.Short
	short.SourceURL = dtoShort.SourceURL
	short.TargetURL = dtoShort.TargetURL
	short.ShortGroup = dtoShort.ShortGroup
	short.IsEnable = dtoShort.IsEnable
	short.ExpireAt = dtoShort.ExpireAt
	short.Remarks = dtoShort.Remarks
	if dtoShort.Length == 0 && (dtoShort.Length < 4 || dtoShort.Length > 16) {
		s.Ctx.WriteString("创建失败，参数错误")
		s.Ctx.ResponseWriter.WriteHeader(400)
		return
	}
	short.FKUser = accInfo.ID
	var err error
	if dtoShort.Automatic {
		err = models.CreateShort(short, dtoShort.Length)
	} else {
		err = models.CreateShortCustom(short)
	}
	if err != nil {
		s.Ctx.WriteString(err.Error())
	}
}

// @Title	UpdateUser
// @Summary	修改一个短链接
// @Param	sid		path	models.Short.ID	true	"短链接id"
// @Param	short	body	models.AddEditShort	true	"body for short"
// @Success	200	{string}	"update success"
// @Failure	403	{string}	"Insufficient user permissions"
// @router /:sid [put]
func (s *ShortController) UpdateShort() {
	var dtoShort models.AddEditShort
	json.Unmarshal(s.Ctx.Input.RequestBody, &dtoShort)

	var short models.Short
	short.ID = s.GetString(":sid")
	accInfo := s.analysisAccountClaims()
	if accInfo.Role == 1 || accInfo.ID == short.FKUser {
		existShort := models.QueryShortByID(short.ID)
		existShort.TargetURL = dtoShort.TargetURL
		existShort.ShortGroup = dtoShort.ShortGroup
		existShort.IsEnable = dtoShort.IsEnable
		existShort.ExpireAt = dtoShort.ExpireAt
		existShort.Remarks = dtoShort.Remarks
		err := models.UpdateShort(short)
		if err != nil {
			s.Ctx.WriteString(err.Error())
			return
		}
		s.Ctx.WriteString("修改成功")
	} else {
		s.Ctx.ResponseWriter.WriteHeader(403)
		s.Ctx.WriteString("无权修改其他用户的链接")
	}
}

// @Title	DeleteShort
// @Summary	根据id删除一个短链接
// @Param	sid	path	string true	"链接id"
// @Success	200	{string}	"delete success!"
// @Failure	403	{string}	"无权删除"
// @router /:sid [delete]
func (s *ShortController) DeleteShort() {
	sid := s.GetString(":sid")
	accInfo := s.analysisAccountClaims()
	short := models.QueryShortByID(sid)
	if accInfo.ID != short.FKUser {
		s.Ctx.ResponseWriter.WriteHeader(403)
		s.Ctx.WriteString("无权删除其他用户的链接")
		return
	}
	if models.DeletedShortUrlById(sid) {
		s.Ctx.WriteString("delete success!")
	} else {
		s.Ctx.WriteString("delete fail!")
	}
}

// @Title	getShortsPage
// @Summary	分頁查詢
// @Date	2022-11-14
// @Auth	sfhj
// @Param	offset	query	int	true	偏移量
// @Param	limit	query	int	true	指定返回记录的数量
// @Param	sort	query	string	true	排序
// @Param	source_url	query	string	false	源url
// @Param	target_url	query	string	false	目标url
// @Param	group	query	string	false	分组
// @Param	is_enable	query	string	false	是否启用
// @Param	exp	query	string	false	过期时间
// @Param	crt	query	string	false	创建时间
// @Param	upt	query	string	false	修改时间
// @Param	det	query	string	false	删除时间
// @Success 200
// @router / [get]
func (s *ShortController) GetShortsPage() {
	accInfo := s.analysisAccountClaims()
	var err error
	var page models.Page
	page.Offset, err = s.GetInt("offset")
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString("请求参数类型错误")
		return
	}
	page.Lmit, err = s.GetInt("limit")
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString("请求参数类型错误")
		return
	}
	page.Sort = analysisOrderBy(s.GetString("sort"))
	fkUser := strconv.FormatUint(uint64(accInfo.ID), 10)
	sourceURL := s.GetString("source_url")
	targetURL := s.GetString("target_url")
	shortGroup := s.GetString("group")
	isEnable := s.GetString("is_enable")
	exp := s.GetString("exp")
	crt := s.GetString("crt")
	upt := s.GetString("upt")
	del := s.GetString("del")
	result, count, err := models.QueryShortsPage(page, fkUser, sourceURL, targetURL, shortGroup, isEnable, exp, crt, upt, del)
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString(err.Error())
		return
	}
	s.Data["json"] = map[string]interface{}{
		"count": count,
		"data":  result,
	}
	s.ServeJSON()
}
