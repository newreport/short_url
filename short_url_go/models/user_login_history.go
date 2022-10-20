package models

import "gorm.io/gorm"

type UserLoginHistory struct {
	gorm.Model
	Ip      string `gorm:"not null"` //当前ip
	FkUser  uint   `gorm:"not null"` //外键关联用户
	MacUuid string `gorm:"not null"` //浏览器的uuid，本地存储长期有效，除非删除
	Browser string `gorm:"not null"` //浏览器
	OS      string `gorm:"not null"` //操作系统
}
