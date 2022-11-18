package controllers

import (
	"encoding/json"
	"fmt"
	"short_url_go/common"
	"short_url_go/models"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

// Operations about Users
type UserController struct {
	BaseController
}

//https://cloud.tencent.com/developer/article/1557075

// @Title users
// @Summary 获取所有用户
// @Description logs.Info user into the system
// @Success 200 {object} models.User
// @Failure 403 User not exist
// @router /all [get]
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
// @Failure 403 Prohibiting the registration
// @router /register [post]
func (u *UserController) Register() {
	// Infos(u)
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Role = 0
	user.DefaultURLLength = 6
	u.Data["json"] = models.CreateUser(user)
	u.ServeJSON()
}

// @Title Login
// @Description logs.Info user into the system
// @Summary 登录
// @Param	body		body 	models.User	true		"The username for login"
// @Success 200 {models.User} Login success
// @Failure 401 The user does not exist.
// @router /login [post]
func (u *UserController) Login() {
	u.infos()
	requestBody := u.JsonData()
	username := requestBody["name"].(string)
	password := requestBody["pwd"].(string)
	user := models.Login(username, password)
	if user.ID > 0 && len(user.Name) > 0 {
		u.Data["json"] = generateRefreshJWT(user.ID)
	} else {
		u.Ctx.ResponseWriter.WriteHeader(401)
		u.Data["json"] = "没有找到用户"
	}
	u.ServeJSON()
}

// @Title account tocken
// @Description logs.Info user into the system
// @Summary 刷新 account tocken
// @Param	jwt		body 	string	true	"The refresh jwt tocken"
// @Success 200 string Refresh success
// @Failure 401 refresh token 失效
// @router /tocken/account [post]
func (u *UserController) RefreshTocken() {
	tokenString := string(u.Ctx.Input.RequestBody)
	token, err := jwt.ParseWithClaims(tokenString, &RefreshClaims{}, func(token *jwt.Token) (interface{}, error) {
		key, _ := common.INIconf.String("JWT::RefreshTokenKey")
		return []byte(key), nil
	})
	if err != nil {
		panic(err)
	}

	if claims, ok := token.Claims.(*RefreshClaims); ok && token.Valid {
		fmt.Println(claims.ID)
		user := models.QueryUserById(claims.ID)
		u.Data["json"] = generateAccountJWT(user)
	} else {
		u.Ctx.ResponseWriter.WriteHeader(401)
		u.Data["json"] = "refresh token 失效，请重新登录"
	}
	u.ServeJSON()
}

// @Title user
// @Summary 新增一个用户
// @Description logs.Info user into the system
// @Param	body	body 	models.User	true	"body for user"
// @Success 200 {bool} Create success
// @Failure 403	Insufficient user permissions
// @router / [post]
func (u *UserController) CreateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	accInfo := u.analysisAccountClaims()
	if accInfo.Role == 1 || accInfo.ID == user.ID {
		u.Data["json"] = models.CreateUser(user)
	} else {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Data["json"] = "无权创建用户"
	}
	u.ServeJSON()
}

// @Title Delete
// @Summary 删除一个用户
// @Description delete the user
// @Param	uid		path 	unit	true	"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 Insufficient user permissions
// @router /:uid [delete]
func (u *UserController) DeleteUser() {
	uid, err := u.GetUint64(":uid")
	if err != nil {
		fmt.Println(err)
	}
	accInfo := u.analysisAccountClaims()
	if accInfo.Role == 1 || accInfo.ID == uint(uid) {
		models.DeleteUser(uint(uid))
		u.Data["json"] = "delete success!"
	} else {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Data["json"] = "无权删除用户"
	}
	u.ServeJSON()
}

// @Summary 修改一个用户
// @Description update the user
// @Param	user	body 	models.User true	"body for user"
// @Success 200 {bool} update success!
// @Failure 403 Insufficient user permissions
// @router / [put]
func (u *UserController) UpdateUser() {
	var user models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	accInfo := u.analysisAccountClaims()
	if accInfo.Role == 1 || accInfo.ID == user.ID {
		u.Data["json"] = models.UpdateUser(user)
	} else {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Data["json"] = "无权修改用户"
	}
	u.ServeJSON()
}

// @Title Update password
// @Summary 修改一个用户的密码
// @Description update the user's password
// @Param	body	body 	string true "body for password"
// @Success 200 {bool} update password success!
// @Failure 403 Insufficient user permissions
// @router /pwd/:uid [patch]
func (u *UserController) UpdateUserPassword() {
	uid := u.GetString(":uid")
	id, _ := strconv.ParseUint(uid, 10, 64)
	var pwd string
	accInfo := u.analysisAccountClaims()
	if uint(id) == accInfo.ID {
		json.Unmarshal(u.Ctx.Input.RequestBody, &pwd)
		u.Data["json"] = models.UpdatePassword(uint(id), pwd)
	} else {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Data["json"] = "无权修改密码"
	}
	u.ServeJSON()
}

// @Title	GetUsersByPage
// @Summary	user分页查询
// @Date	2022-11-18
// @Auth	sfhj
// @Param	page	query	models.Page	true	分页
// @Param	query	query	models.UserPageQuery	false	查询参数
// @Success	200
// @Failure 403 Insufficient user permissions
// @router / [get]
func (u *UserController) GetUsersByPage() {
	accInfo := u.analysisAccountClaims()
	if accInfo.Role == 0 {
		u.Ctx.ResponseWriter.WriteHeader(403)
		u.Data["json"] = "无权查询其他用户"
	}
	var err error
	var page models.Page
	page.Offset, err = u.GetInt("offset")
	if err != nil {
		u.Ctx.ResponseWriter.WriteHeader(400)
		u.Data["json"] = "请求参数类型错误"
	}
	page.Lmit, err = u.GetInt("limit")
	if err != nil {
		u.Ctx.ResponseWriter.WriteHeader(400)
		u.Data["json"] = "请求参数类型错误"
	}
	page.Sort = analysisOrderBy(u.GetString("sort"))
	var query models.UserPageQuery
	query.Name = u.GetString("name")
	query.Nickname = u.GetString("nickname")
	query.Group = u.GetString("group")
	query.Role = u.GetString("role")
	query.CreatedAt = u.GetString("crt")
	query.UpdatedAt = u.GetString("upt")
	query.DeletedAt = u.GetString("det")
	result, count, err := models.QueryPageUsers(query, page)
	u.Data["json"] = map[string]interface{}{
		"count": count,
		"data":  result,
	}
	if err != nil {
		u.Ctx.ResponseWriter.WriteHeader(400)
		u.Data["json"] = "请求参数错误"
	}
	u.ServeJSON()

}
