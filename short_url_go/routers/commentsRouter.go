package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "CreateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUser",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "DeleteUser",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAllUsers",
            Router: `/all`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "UpdateUserPassword",
            Router: `/pwd/:uid`,
            AllowHTTPMethods: []string{"patch"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["short_url_go/controllers:UserController"] = append(beego.GlobalControllerRouter["short_url_go/controllers:UserController"],
        beego.ControllerComments{
            Method: "RefreshTocken",
            Router: `/tocken/account`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
