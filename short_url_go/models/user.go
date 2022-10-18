package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name             string `gorm:"not null"` //用户名，登录名称
	NickName         string `gorm:"not null"`
	Passwd           string `gorm:"not null"`
	Role             int8   `gorm:"not null"`
	DefaultUrlLength uint8  `gorm:"not null"`
	Remarks          string
	baseSqlStruct
}

type UserInterface interface {
	BaseSqlInterface[User]
	SayUser()
}

func (T User) SayUser() {
	fmt.Println("this is User")
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
