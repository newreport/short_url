package models

import "gorm.io/gorm"

type ShortGroup struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`   //分组名称
	FkUser uint   `json:"fkUser" gorm:"not null"` //外键用户
}
