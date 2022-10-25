package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/core/config"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID               uint `gorm:"primaryKey;<-:create"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt `gorm:"index"`
	Name             string         `gorm:"not null"` //用户名，登录名称
	Nickname         string         `gorm:"not null"`
	Password         string         `gorm:"not null"`
	Role             int8           `gorm:"not null"`
	DefaultUrlLength uint8          `gorm:"not null"`
	Group            string
	Remarks          string
}

type UserShort string

const (
	Id        UserShort = "id"
	Name                = "name"
	NickName            = "nickname"
	CreatedAt           = "created_at"
	Group               = "group"
)

var U5Seed uuid.UUID

func init() {
	ini, err := config.NewConfig("ini", "conf/secret.conf")
	if err != nil {
		panic(err)
	}
	pwdUUID, err := ini.String("UUID::UserPwd")
	if err != nil {
		fmt.Println(err)
	}
	U5Seed = uuid.Must(uuid.FromString(pwdUUID))
	fmt.Println("U5Seed:", U5Seed)
}

func GetAllUsers() []User {
	var users []User
	DB.Order("created_at desc").Find(&users)
	return users
}

func Login(username, password string) User {
	var user User
	password = uuid.NewV5(U5Seed, password).String()
	DB.Model(&User{}).Where(&User{Name: username, Password: password}).First(&user)
	return user
}

func CreateUser(user User) uint {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	DB.Create(&user)
	return user.ID
}

func DeleteUser(id uint) bool {
	var count int64
	DB.Model(Short{}).Where("fk_user = ? ", id).Count(&count)
	if count > 0 {
		return false
	}
	result := DB.Delete(&User{}, id)
	return result.RowsAffected > 0
}

func UpdateUser(user User) bool {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	result := DB.Model(&user).Updates(User{Name: user.Name, Nickname: user.Nickname, Password: user.Password, Role: user.Role, DefaultUrlLength: user.DefaultUrlLength})
	return result.RowsAffected > 0
}

// @Title	分页查询
// @Auth	sfhj
// @Date	2022-10-25
// @Param	name	string 用户名
// @Param	nicknmae	string 昵称
// @Param	group	string 分组
// @Param	role	uint	权限
// @Param	page	models.Page	分页查询累
// @Return	users	[]models.User
func QueryPageUsers(name string, nicknmae string, group string, role uint, page Page) []User {
	express := DB.Model(&User{})
	if len(name) > 0 {
		express = express.Where("name LIKE % ? % ", name)
	}
	if len(nicknmae) > 0 {
		express = express.Where("nickname LIKE % ? % ", nicknmae)
	}
	if role == 0 || role == 1 {
		express = express.Where("role = ? ", nicknmae)
	}
	if len(group) > 0 {
		express = express.Where("group = ? ", group)
	}
	express = express.Limit(page.PageSize).Offset((page.PageNum - 1) * page.PageSize)
	if page.Desc {
		express = express.Order(page.Keyword + " desc")
	} else {
		express = express.Order(page.Keyword)
	}
	var result []User
	express.Select("id", "created_at", "updated_at", "name", "nickname", "role", "default_url_length", "group", "remarks").Find(&result)
	return result
}

func UpdatePassword(id uint, pwd string) bool {
	pwd = uuid.NewV5(U5Seed, pwd).String()
	result := DB.Model(&User{}).Where("id = ?", id).Update("password", pwd)
	return result.RowsAffected > 0
}

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
