package controllers

import (
	"encoding/json"
	"short_url_go/common"
	"strings"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/golang-jwt/jwt/v4"
)

//http://bartontang.github.io/2017/10/12/beego-%E8%8E%B7%E5%8F%96request-body%E5%86%85%E5%AE%B9/

type BaseController struct {
	beego.Controller
}

func (b *BaseController) RequestBody() []byte {
	return b.Ctx.Input.RequestBody
}
func (b *BaseController) decodeRawRequestBodyJson() map[string]interface{} {
	var mm interface{}
	requestBody := make(map[string]interface{})
	json.Unmarshal(b.RequestBody(), &mm)
	if mm != nil {
		var m1 map[string]interface{}
		m1 = mm.(map[string]interface{})
		for k, v := range m1 {
			requestBody[k] = v
		}
	}
	return requestBody
}

func (b *BaseController) JsonData() map[string]interface{} {
	return b.decodeRawRequestBodyJson()
}

// @Title analysisAccountClaims
// @Description 從 http head 中解析 Account Token
// @Param b beego.Controller	http 上下文
// @Return result  controllers.AccountClaims
func (b *BaseController) analysisAccountClaims() (result AccountClaims) {
	tokenString := b.Ctx.Input.Header("Authorization")
	tokenString = strings.Split(tokenString, "")[1]
	token, _ := jwt.ParseWithClaims(tokenString, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
		key, _ := common.INIconf.String("JWT::AccessTokenKey")
		return []byte(key), nil
	})
	if claims, ok := token.Claims.(*AccountClaims); ok && token.Valid {
		result = *claims
	}
	return
}

// @Title analysisOrderBy
// @Description 计算orderby
func analysisOrderBy(str string) string {
	str = strings.Replace(str, "+", " asc", -1)
	str = strings.Replace(str, "-", " desc", -1)
	return str
}
