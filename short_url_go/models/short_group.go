package models

import "gorm.io/gorm"

type ShortGroup struct {
	gorm.Model
	Name   string `gorm:"not null"`
	FkUser uint   `gorm:"not null"` //外键
}
