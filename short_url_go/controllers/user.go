package controllers

import (
	"encoding/json"
	"fmt"
	"short_url_go/models"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description logs.Info user into the system
// @Summary 登录
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Println("登录中...", password)

	user := models.Login(username, password)
	if user.ID > 0 && len(user.Name) > 0 {
		u.Data["json"] = user
	} else {
		u.Data["json"] = "user not exist"
	}
	u.ServeJSON()
}

//https://cloud.tencent.com/developer/article/1557075

// @Title users
// @Summary 获取所有用户
// @Description logs.Info user into the system
// @Success 200 {object} models.User
// @Failure 403 user not exist
// @router / [get]
func (u *UserController) GetAllUsers() {
	u.Data["json"] = models.GetAllUsers()
	// var num uint
	// var num2 int
	// var num3 int8
	// var num4 uint64
	// fmt.Println("uint:", unsafe.Sizeof(num))
	// fmt.Println("uint64:", unsafe.Sizeof(num4))
	// fmt.Println("int:", unsafe.Sizeof(num2))
	// fmt.Println("int8:", unsafe.Sizeof(num3))
	u.ServeJSON()
}

// @Title Register
// @Summary 注册

// @Description logs.Info user into the system
// @Param	body		body 	models.User	true	"body for user"
// @Success 200 {bool} register success
// @Failure 403 not role
// @router /register [post]
func (u *UserController) Register() {
	// Infos(u)
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Role = 0
	user.DefaultUrlLength = 6
	u.Data["json"] = models.CreateUser(user)
	u.ServeJSON()
}

// @Title user
// @Summary 新增一个用户
// @Description logs.Info user into the system
// @Param	body	body 	models.User	true	"body for user"
// @Success 200 {bool} create success
// @Failure 403 not role
// @router / [post]
func (u *UserController) CreateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Name = "asd"
	u.Data["json"] = models.CreateUser(user)
	u.ServeJSON()
}

// @Title Delete
// @Summary 删除一个用户
// @Description delete the user
// @Param	uid		path 	unit	true	"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) DeleteUser() {
	uid, err := u.GetUint64(":uid")
	if err != nil {
		fmt.Println(err)
	}
	models.DeleteUser(uint(uid))
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title update
// @Summary 修改一个用户
// @Description update the user
// @Param	user	body 	models.User true	"body for user"
// @Success 200 {bool} update success!
// @Failure 403 not have role
// @router / [put]
func (u *UserController) UpdateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	u.Data["json"] = models.CreateUser(user)
	u.ServeJSON()
}

// @Title update_password
// @Summary 修改一个用户的密码
// @Description update the user's password
// @Param	password	body 	models.Page true "body for string"
// @Success 200 {bool} update success!
// @Failure 403 not have role
// @router /pwd/:uid [patch]
func (u *UserController) UpdateUserPassword() {

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
	// logs.Info(u.Ctx.Input.RequestBody) //该函数返回 Get 请求和 Post 请求中的所有数据，和 PHP 中$_REQUEST 类似

}
