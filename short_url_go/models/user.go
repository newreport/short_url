package models

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	NickName string `gorm:"not null"`
	Passwd   string `gorm:"not null"`
	Role     int8   `gorm:"not null"`
	Remarks  string
	BaseStruct
}

type UserMethod interface {
	baseSqlMethod[User]
}

func (T User) SayUser() {
	fmt.Println("this is User")
}
func Test() {
	var u UserMethod
	var us User
	u = us
	u.FirstOrDefault()
	// fmt.Println(us.FirstOrDefault().Name)
	// fmt.Println(FirstOrDefault[User]().Name)

}
