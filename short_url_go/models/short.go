package models

import (
	"time"

	"gorm.io/gorm"
)

type Short struct {
	Sid       string `gorm:"primaryKey,size:50;"`
	SourceUrl string `gorm:"not null"`
	TargetUrl string `gorm:"not null"`
	Remarks   string
	FkUser    uint `gorm:"not null"` //外键
	UrlGroup  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// 添加一条URL短链接
func AddOneUrlDefault(url string, userId int) {

}

// 添加一条指定长度的短链接
func AddOneUrlAssignLength(url string, userId int, lengthNum int) {

}

// 添加一条自定义长度的短链接
func AddOneUrl(sourceUrl string, targetUrl string) {

}
