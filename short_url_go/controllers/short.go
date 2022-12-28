package controllers

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"html/template"
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

// @Title	UpdateShort
// @Summary	修改一个短链接
// @Param	id		path	models.Short.ID	true	"短链接id"
// @Param	short	body	models.AddEditShort	true	"body for short"
// @Success	200	{string}	"update success"
// @Failure	403	{string}	"Insufficient user permissions"
// @router /:id [put]
func (s *ShortController) UpdateShort() {
	var dtoShort models.AddEditShort
	json.Unmarshal(s.Ctx.Input.RequestBody, &dtoShort)

	var short models.Short
	short.ID = s.GetString(":id")
	accInfo := s.analysisAccountClaims()
	if accInfo.Role == 1 || accInfo.ID == short.FKUser {
		existShort := models.QueryShortByID(short.ID)
		existShort.TargetURL = dtoShort.TargetURL
		existShort.ShortGroup = dtoShort.ShortGroup
		existShort.IsEnable = dtoShort.IsEnable
		existShort.ExpireAt = dtoShort.ExpireAt
		existShort.Remarks = dtoShort.Remarks
		//自动生成
		if dtoShort.Automatic {
			existShort.TargetURL = models.GenerateUrl(short.SourceURL, short.FKUser, dtoShort.Length)
		}
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

// @Title	RecoverShort
// @Summary 恢復一個鏈接
// @Param	id		path	models.Short.ID	true	"短链接id"
// @Success	200	{string}	"recover success"
// @Failure	403	{string}	"Insufficient user permissions"
// @router /:id/recover [put]
func (s *ShortController) RecoverShort() {

	id := s.GetString(":id")
	accInfo := s.analysisAccountClaims()
	existShort := models.QueryShortByID(id)
	if accInfo.Role == 1 || accInfo.ID == existShort.FKUser {
		if models.RecoverShort(id) {
			s.Ctx.WriteString("恢復成功")
		} else {
			s.Ctx.WriteString("恢復失敗")
		}
	} else {
		s.Ctx.ResponseWriter.WriteHeader(403)
		s.Ctx.WriteString("无权恢復其他用户的链接")
	}
}

// @Title	ExportHtml
// @Summary	导出html静态页
// @Param	id		path	models.User.ID	true	"短链接id"
// @Param	short	body	models.AddEditShort	true	"body for short"
// @Success	200	{file}	"get success"
// @Failure	403	{string}	"Insufficient user permissions"
// @router /html/:id [get]
func (s *ShortController) ExportHtml() {
	accountInfo := s.analysisAccountClaims()
	t := template.New("html")
	text := `<script>window.location.href="{{.}}"</script>`
	t.Parse(text)
	id, err := s.GetInt(":id")
	if err != nil {
		s.Ctx.ResponseWriter.WriteHeader(400)
		s.Ctx.WriteString("请求参数类型错误")
		return
	}
	if uint(id) != accountInfo.ID {
		s.Ctx.ResponseWriter.WriteHeader(403)
		s.Ctx.WriteString("权限错误")
		return
	}
	data := models.QueryAllByUserID(accountInfo.ID)

	// data := map[string]string{"1.html": "hello", "2.html": "world"}
	buf := new(bytes.Buffer)
	// 创建一个压缩文档
	w := zip.NewWriter(buf)

	for k, v := range data {
		bufHtml := new(bytes.Buffer)
		err = t.Execute(bufHtml, v)
		if err != nil {
			panic(err)
		}
		f, err := w.Create(k + ".html")
		if err != nil {
			panic(err)
		}
		_, err = f.Write(bufHtml.Bytes())
		if err != nil {
			panic(err)
		}
	}
	err = w.Close()
	if err != nil {
		panic(err)
	}

	// f, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0666)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// buf.WriteTo(f)
	// s.Ctx.Output.Download("file.zip")

	s.Ctx.Output.Header("Content-Type", "application/octet-stream")
	s.Ctx.Output.Header("Content-Transfer-Encoding", "binary")
	s.Ctx.Output.Header("Content-Disposition", "attachment; filename=file.zip")
	s.Ctx.Output.Header("Content-Description", "File Transfer")
	s.Ctx.Output.Header("Accept-Ranges", "bytes")
	s.Ctx.Output.Body(buf.Bytes())
}

// @Title	DeleteShort
// @Summary	根据id删除一个短链接
// @Param	id	path	string true	"链接id"
// @Success	200	{string}	"delete success!"
// @Failure	403	{string}	"无权删除"
// @router /:id [delete]
func (s *ShortController) DeleteShort() {
	id := s.GetString(":id")
	accInfo := s.analysisAccountClaims()
	short := models.QueryShortByID(id)
	if accInfo.ID != short.FKUser {
		s.Ctx.ResponseWriter.WriteHeader(403)
		s.Ctx.WriteString("无权删除其他用户的链接")
		return
	}
	if models.DeletedShortUrlByID(id) {
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
// @Param	sort	query	string	false	排序
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
