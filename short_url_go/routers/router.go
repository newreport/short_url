// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"short_url_go/controllers"
	"strings"

	"github.com/beego/beego/v2/server/web/context"

	"github.com/beego/beego/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}

// 过滤器
var FilterToken = func(ctx *context.Context) {
	logs.Info("current router path is ", ctx.Request.RequestURI)
	if ctx.Request.RequestURI != "/login" && ctx.Request.RequestURI != "/register" && ctx.Input.Header("Authorization") == "" {
		logs.Error("without token, unauthorized !!")
		ctx.ResponseWriter.WriteHeader(401)
		ctx.ResponseWriter.Write([]byte("no permission"))
	}
	if ctx.Request.RequestURI != "/login" && ctx.Input.Header("Authorization") != "" {
		token := ctx.Input.Header("Authorization")
		token = strings.Split(token, "")[1]
		logs.Info("curernttoken: ", token)
		// validate token
		// invoke ValidateToken in utils/token
		// invalid or expired todo res 401
	}
}
