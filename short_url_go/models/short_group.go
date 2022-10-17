package models

import "gorm.io/gorm"

type ShortGroup struct {
	gorm.Model
	Name   string `gorm:"not null"` //分组名称
	FkUser uint   `gorm:"not null"` //外键用户
}
