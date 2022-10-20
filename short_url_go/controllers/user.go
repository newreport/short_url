package controllers

import (
	"fmt"
	"short_url_go/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

//https://cloud.tencent.com/developer/article/1557075

// @Title GetAll
// @Description Logs user into the system
// @Success 200 {string}
// @Failure 403 user not exist
// @router /GetAll [get]
func (u *UserController) AllUsers() {
	u.Data["json"] = models.GetAllUsers()
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
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

// @Title Register
// @Description Logs user into the system
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {bool} register success
// @Failure 403 not role
// @router /register [post]
func (u *UserController) Register() {
	var user models.User
	data := u.Ctx.Input.RequestBody
	fmt.Println(data)
	user.Name = u.GetString("username")
	user.Passwd = u.GetString("password")
	user.NickName = u.GetString("nickname")
	user.Role = 0
	user.DefaultUrlLength = 6
	u.Data["json"] = models.AddUser(user)
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	fmt.Println(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}
