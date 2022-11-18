package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// 用户表
type User struct { //用户表
	ID               uint           `json:"ID" gorm:"primaryKey;<-:create"` //id
	CreatedAt        time.Time      `json:"crt" gorm:"<-:create"`           //创建时间
	UpdatedAt        time.Time      `json:"upt" gorm:"<-"`                  //最后更新时间
	DeletedAt        gorm.DeletedAt `json:"det" gorm:"index"`               //软删除时间
	Name             string         `json:"name" gorm:"not null"`           //用户名，登录名称
	Nickname         string         `json:"nickname" gorm:"not null"`       //昵称
	Password         string         `json:"pwd" gorm:"not null"`            //密码
	Role             int8           `json:"role" gorm:"not null"`           //角色
	DefaultURLLength uint8          `json:"urlLength" gorm:"not null"`      //配置项：url默认长度
	AuthorURL        string         `json:"authorUrl"`                      //头像地址
	Phone            string         `json:"phone"`                          //手机号
	Group            string         `json:"group"`                          //分组
	Remarks          string         `json:"remarks"`                        //备注
	I18n             string         `json:"i18n"`                           //国际化
	AutoInsertSpace  bool           `json:"autoInsertSpace"`                //盘古之白
}

// 分页查询
type UserPageQuery struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Group     string `json:"group"`
	Role      string `json:"role"`
	CreatedAt string `json:"crt"`
	UpdatedAt string `json:"upt"`
	DeletedAt string `json:"det"`
}

// @Title 获取所有用户
func GetAllUsers() []User {
	var users []User
	DB.Order("created_at desc").Find(&users)
	return users
}

// @Title 登录
func Login(username, password string) User {
	var user User
	password = uuid.NewV5(U5Seed, password).String()
	DB.Model(&User{}).Where(&User{Name: username, Password: password}).First(&user)
	return user
}

// @Title 创建用户
func CreateUser(user User) uint {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	DB.Create(&user)
	return user.ID
}

// @Title 根据id删除用户
func DeleteUser(id uint) bool {
	var count int64
	//存在url不允许删除
	DB.Model(Short{}).Where("fk_user = ? ", id).Count(&count)
	if count > 0 {
		return false
	}
	result := DB.Delete(&User{}, id)
	return result.RowsAffected > 0
}

// @Title 更新用户信息，除了id
func UpdateUser(user User) bool {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	result := DB.Model(&user).Updates(User{Name: user.Name, Nickname: user.Nickname, Password: user.Password, Role: user.Role, AuthorURL: user.AuthorURL, Phone: user.Phone, Group: user.Group, I18n: user.I18n, AutoInsertSpace: user.AutoInsertSpace, Remarks: user.Remarks, DefaultURLLength: user.DefaultURLLength})
	return result.RowsAffected > 0
}

// @Title	分页查询
// @Auth	sfhj
// @Date	2022-11-15
// @Param	query	models.UserPageQuery	查询参数
// @Param	page	models.Page	分页查询struct
// @Return	users	[]models.User,error
func QueryPageUsers(query UserPageQuery, page Page) (result []User, count int64, err error) {
	express := DB.Model(&User{})
	if analysisRestfulRHS(express, "name", query.Name) &&
		analysisRestfulRHS(express, "nickname", query.Nickname) &&
		analysisRestfulRHS(express, "role", query.Nickname) {
		express.Count(&count)
		express = express.Order(page.Sort).Limit(page.Lmit).Offset((page.Offset - 1) * page.Lmit)
		express.Select("id", "created_at", "updated_at", "name", "nickname", "role", "default_url_length", "group", "remarks").Find(&result)
	} else {
		err = errors.New("查詢參數錯誤")
	}
	return
}

// @Title 根据id查询用户
func QueryUserById(id uint) User {
	var user User
	DB.First(&user, id)
	return user
}

// @Title 根据用户id修改密码
func UpdatePassword(id uint, pwd string) bool {
	pwd = uuid.NewV5(U5Seed, pwd).String()
	result := DB.Model(&User{}).Where("id = ?", id).Update("password", pwd)
	return result.RowsAffected > 0
}

// @Title 根据用户id集合删除用户
func DeleteUsers(ids []uint) bool {
	var count int64
	DB.Model(Short{}).Where("fk_user IN ?", ids).Count(&count)
	if count > 0 {
		return false
	}
	var users []User
	result := DB.Model(User{}).Where(&users, ids).Delete(&User{})
	return result.RowsAffected > 0
}
