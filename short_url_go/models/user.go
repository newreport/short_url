package models

import (
	"fmt"
	"short_url_go/common"

	"github.com/beego/beego/v2/core/config"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"not null"` //用户名，登录名称
	NickName         string `gorm:"not null"`
	Password         string `gorm:"not null"`
	Role             int8   `gorm:"not null"`
	DefaultUrlLength uint8  `gorm:"not null"`
	Remarks          string
	baseSqlStruct
}

func Test() {

	// u := FirstOrDefault[User]("id = ?", 1)
	// fmt.Println(u)
	// fmt.Println(reflect.TypeOf(map[string]interface{}{"name": "jinzhu", "age": 20}))
	// u2 := Where[User]([]int64{1})
	// fmt.Println(u2)
	// s := Where[Short]()
	// fmt.Println(s)
	// fmt.Print("MD5:")
	// str := common.MD5("baidu.com")
	// fmt.Println(str)
	// GenerateUrlDefault("baidu.com")

	// var u BaseSqlInterface[User]
	// var us User
	// u = us
	// u.FirstOrDefault()

	// var u UserInterface
	// var us User
	// u = us
	// u.FirstOrDefault()
	// u.SayUser()
	// var models []Short
	// arr := [2]string{"google.com"}
	// common.DB.Debug().Where(map[string]interface{}{"source_url": arr}).Find(&models)
	// fmt.Println(models)
}

type UserInterface interface {
	BaseSqlInterface[User]
	SayUser()
}

func (T User) SayUser() {
	fmt.Println("this is User")
}

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
	fmt.Println("u5Ini:", U5Seed)
}

func GetAllUsers() []User {
	var users []User
	common.DB.Find(&users)
	return users
}

func Login(username, password string) User {
	var user User
	fmt.Println("u5:", U5Seed)
	fmt.Println("be", uuid.NewV5(U5Seed, password).String())
	password = uuid.NewV5(U5Seed, password).String()
	fmt.Print("login:")
	fmt.Println(password)
	common.DB.Model(&User{}).Where(&User{Name: username, Password: password}).First(&user)
	return user
}

func CreateUser(user User) bool {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	result := common.DB.Create(&user)
	return result.RowsAffected > 0
}

func DeleteUser(id uint) bool {
	var count int64
	common.DB.Model(Short{}).Where("fk_user = ? ", id).Count(&count)
	if count > 0 {
		return false
	}
	result := common.DB.Delete(&User{}, id)
	return result.RowsAffected > 0
}

func UpdateUser(user User) bool {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	result := common.DB.Model(&user).Updates(User{Name: user.Name, NickName: user.NickName, Password: user.Password, Role: user.Role, DefaultUrlLength: user.DefaultUrlLength})
	return result.RowsAffected > 0
}

func UpdatePassword(id uint, pwd string) bool {
	pwd = uuid.NewV5(U5Seed, pwd).String()
	result := common.DB.Model(&User{}).Where("id = ?", id).Update("password", pwd)
	return result.RowsAffected > 0
}

func DeleteUsers(ids []uint) bool {
	var count int64
	common.DB.Model(Short{}).Where("fk_user IN ?", ids).Count(&count)
	if count > 0 {
		return false
	}
	var users []User
	result := common.DB.Model(User{}).Where(&users, ids).Delete(&User{})
	return result.RowsAffected > 0
}
