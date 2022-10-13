package models

import (
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	NickName string `gorm:"not null"`
	Passwd   string `gorm:"not null"`
	Role     int8   `gorm:"not null"`
	Remarks  string
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
	fmt.Println(reflect.TypeOf(map[string]interface{}{"name": "jinzhu", "age": 20}))
	u2 := Where[User]([]int64{1})
	fmt.Println(u2)

	// var u BaseSqlInterface[User]
	// var us User
	// u = us
	// u.FirstOrDefault()

	// var u UserInterface
	// var us User
	// u = us
	// u.FirstOrDefault()
	// u.SayUser()
}
