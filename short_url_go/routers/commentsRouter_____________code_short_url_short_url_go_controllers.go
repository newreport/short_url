package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["short_url_go/controllers:LoginController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: "/logout",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:LoginController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout1",
            Router: "/logout1",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
