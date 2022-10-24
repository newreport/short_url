package controllers

import (
	"encoding/json"
	"fmt"
	"short_url_go/models"
	"unsafe"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Login
// @Description Logs user into the system
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
// @Description Logs user into the system
// @Success 200 {object} models.User
// @Failure 403 user not exist
// @router / [get]
func (u *UserController) GetAllUsers() {
	u.Data["json"] = models.GetAllUsers()
	var num uint
	var num2 int
	var num3 int8
	var num4 uint64
	fmt.Println("uint:", unsafe.Sizeof(num))
	fmt.Println("uint64:", unsafe.Sizeof(num4))
	fmt.Println("int:", unsafe.Sizeof(num2))
	fmt.Println("int8:", unsafe.Sizeof(num3))
	u.ServeJSON()
}

// @Title register
// @Summary 注册
// @Description Logs user into the system
// @Param	register		body 	models.User	true	"body for user"
// @Success 200 {bool} register success
// @Failure 403 not role
// @router /register [post]
func (u *UserController) Register() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Role = 0
	user.DefaultUrlLength = 6
	u.Data["json"] = models.CreateUser(user)
	u.ServeJSON()
}

// @Title user
// @Summary 新增一个用户
// @Description Logs user into the system
// @Param	body	body 	models.User	true	"body for user"
// @Success 200 {bool} create success
// @Failure 403 not role
// @router / [post]
func (u *UserController) CreateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
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
// @Param	password	body 	string true "body for string"
// @Success 200 {bool} update success!
// @Failure 403 not have role
// @router /pwd/:uid [patch]
func (u *UserController) UpdateUserPassword() {

}
