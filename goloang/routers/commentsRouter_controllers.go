package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["goloang/controllers:LoginController"] = append(beego.GlobalControllerRouter["goloang/controllers:LoginController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: "/getall",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goloang/controllers:ShortController"] = append(beego.GlobalControllerRouter["goloang/controllers:ShortController"],
        beego.ControllerComments{
            Method: "AddOne",
            Router: "/:url",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goloang/controllers:ShortController"] = append(beego.GlobalControllerRouter["goloang/controllers:ShortController"],
        beego.ControllerComments{
            Method: "EditOne",
            Router: "/editone",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["goloang/controllers:ShortController"] = append(beego.GlobalControllerRouter["goloang/controllers:ShortController"],
        beego.ControllerComments{
            Method: "GetList",
            Router: "/getlist",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
