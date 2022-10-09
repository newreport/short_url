package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not null"`
	NickName string `gorm:"not null"`
	Passwd   string `gorm:"not null"`
	Role     int8   `gorm:"not null"`
	Remarks  string
}
