package controllers

import (
	"encoding/json"
	"fmt"
	"short_url_go/models"

	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type UserController struct {
	beego.Controller
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

// @Title register
// @Description Logs user into the system
// @Param	register		body 	models.User		"body for user"
// @Success 200 {bool} register success
// @Failure 403 not role
// @router /register [post]
func (u *UserController) Register() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Role = 0
	user.DefaultUrlLength = 6
	u.Data["json"] = models.AddUser(user)
	u.ServeJSON()
}

// // @Title user
// // @Description Logs user into the system
// // @Param	user		body 	models.User		"body for user"
// // @Success 200 {bool} register success
// // @Failure 403 not role
// // @router / [post]
// func (u *UserController) AddUser() {
// 	var user models.User
// 	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
// 	u.Data["json"] = models.AddUser(user)
// 	u.ServeJSON()
// }

//https://cloud.tencent.com/developer/article/1557075

// @Title users
// @Description Logs user into the system
// @Success 200 {string}
// @Failure 403 user not exist
// @router / [get]
func (u *UserController) GetAllUsers() {
	u.Data["json"] = models.GetAllUsers()
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
