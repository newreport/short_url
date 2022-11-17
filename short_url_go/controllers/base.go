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



// @Title analysisOrderBy
// @Description 计算orderby
func analysisOrderBy(str string) string {
	str = strings.Replace(str, "+", " asc", -1)
	str = strings.Replace(str, "-", " desc", -1)
	return str
}



func Infos(u *UserController) {
	logs.Info(u.Ctx.Input.Protocol())    //获取用户请求的协议:HTTP/1.1
	logs.Info(u.Ctx.Input.URI())         //用户请求的RequestURI: /v1/apptodayRpt/UpALL
	logs.Info(u.Ctx.Input.URL())         //请求的URL地址: /v1/apptodayRpt/UpALL
	logs.Info(u.Ctx.Input.Scheme())      //请求的 scheme: http/https
	logs.Info(u.Ctx.Input.Domain())      //请求的域名:例如 beego.me, 192.168.0.120
	logs.Info(u.Ctx.Input.SubDomains())  //返回请求域名的根域名,例如请求是blog.beego.me-->返回 beego.me;192.168.0.120--> 192.168
	logs.Info(u.Ctx.Input.Host())        //请求的域名,和上面相同:例如 beego.me, 192.168.0.120
	logs.Info(u.Ctx.Input.Site())        //请求的站点地址,scheme+doamin的组合: http://192.168.0.10
	logs.Info(u.Ctx.Input.Method())      //请求的方法:GET,POST 等
	logs.Info(u.Ctx.Input.Is("POST"))    //判断是否是某一个方法:是不是POST方法,注意必须大写
	logs.Info(u.Ctx.Input.IsGet())       //是不是Get请求
	logs.Info(u.Ctx.Input.IsPut())       //是不是Put请求
	logs.Info(u.Ctx.Input.IsPost())      //是不是Post请求
	logs.Info(u.Ctx.Input.IsAjax())      //判断是否是AJAX请求:false
	logs.Info(u.Ctx.Input.IsSecure())    //判断当前请求是否HTTPS请求:false
	logs.Info(u.Ctx.Input.IsWebsocket()) //判断当前请求是否 Websocket请求:false
	logs.Info(u.Ctx.Input.IsUpload())    //判断当前请求是否有文件上传:true
	logs.Info(u.Ctx.Input.IP())          //返回请求用户的 IP,如果用户通过代理，一层一层剥离获取真实的IP:192.168.0.102
	logs.Info(u.Ctx.Input.Proxy())       //返回用户代理请求的所有IP,如果没有代理,返回[]
	logs.Info(u.Ctx.Input.Port())        //返回请求的服务器端口:3000
	logs.Info(u.Ctx.Input.UserAgent())   //客户端浏览器的信息:Mozilla/5.0 (Linux; Android 5.1.1; vivo X7 Build/LMY47V) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/39.0.0.0 Mobile Safari/537.36 Html5Plus/1.0 (Immersed/24.0)
	logs.Info(u.Ctx.Input.Query("name")) //该函数返回 Get 请求和 Post 请求中的所有数据，和 PHP 中$_REQUEST 类似
}
