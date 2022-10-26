package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
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
