﻿package models

import (
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type User struct {
	ID               uint      `json:"id" gorm:"primaryKey;<-:create"`       //id
	CreatedAt        time.Time `json:"crt" gorm:"<-:create"`                 //创建时间
	UpdatedAt        time.Time `json:"upt" gorm:"<-"`                        //最后更新时间
	Name             string    `json:"name" gorm:"not null"`                 //用户名，登录名称
	Nickname         string    `json:"nickname" gorm:"not null;uniqueIndex"` //昵称
	Password         string    `json:"pwd" gorm:"not null"`                  //密码
	Role             int8      `json:"role" gorm:"not null"`                 //角色
	DefaultURLLength uint8     `json:"urlLength" gorm:"not null"`            //配置项：url默认长度
	Author           []byte    `json:"author"`                               //头像地址
	Phone            string    `json:"phone"`                                //手机号
	Group            string    `json:"group"`                                //分组
	Remarks          string    `json:"remarks"`                              //备注
	I18n             string    `json:"i18n"`                                 //国际化
	AutoInsertSpace  bool      `json:"autoInsertSpace"`                      //盘古之白
	Domain           string    `json:"domain" gorm:"uniqueIndex"`            //域名
}

func Login(username, password string) User {
	var user User
	password = uuid.NewV5(U5Seed, password).String()
	DB.Model(&User{}).Where(&User{Name: username, Password: password}).First(&user)
	return user
}

func CreateUser(user User) uint {
	user.Password = uuid.NewV5(U5Seed, user.Password).String()
	DB.Create(&user)
	return user.ID
}

func DeleteUser(id uint) bool {
	DB.Unscoped().Delete(&Short{}, "fk_user = ?", id)
	result := DB.Delete(&User{}, id)
	return result.RowsAffected > 0
}

func DeleteUsers(ids []uint) bool {
	DB.Delete(&Short{}, "fk_user IN ?", ids)
	result := DB.Delete(&User{}, "id IN ? ", ids)
	return result.RowsAffected > 0
}

func UpdateUser(user User) bool {
	var existUser User
	if DB.Unscoped().Where("(name = ? OR domain = ?) AND id != ?", user.Name, user.Domain, user.ID).First(&existUser).RowsAffected > 0 {
		return false
	}
	if existUser.Password != user.Password {
		user.Password = uuid.NewV5(U5Seed, user.Password).String()
	}
	result := DB.Model(&user).Updates(User{Name: user.Name, Nickname: user.Nickname, Password: user.Password, Role: user.Role, Author: user.Author, Phone: user.Phone, Group: user.Group, I18n: user.I18n, AutoInsertSpace: user.AutoInsertSpace, Remarks: user.Remarks, DefaultURLLength: user.DefaultURLLength, Domain: user.Domain})
	return result.RowsAffected > 0
}

func QueryUsersPage(page Page, name string, nickname string, role string, group string, phone string, domain string) (result []User, count int64, err error) {
	express := DB.Model(&User{})
	if analysisRestfulRHS(express, "name", name) &&
		analysisRestfulRHS(express, "nickname", nickname) &&
		analysisRestfulRHS(express, "role", role) &&
		analysisRestfulRHS(express, "phone", phone) &&
		analysisRestfulRHS(express, "domain", domain) &&
		analysisRestfulRHS(express, "group", group) {
		express.Count(&count)
		express = express.Order(page.Sort).Limit(page.Lmit).Offset((page.Offset - 1) * page.Lmit)
		express.Select("id", "created_at", "updated_at", "name", "nickname", "role", "password", "default_url_length", "group", "i18n", "auto_insert_space", "remarks", "domain", "phone").Find(&result)
	} else {
		err = errors.New("查詢參數錯誤")
	}
	return
}

func QueryUserByID(id uint) User {
	var user User
	DB.Unscoped().First(&user, id)
	return user
}

func QueryUserByDomain(domain string) User {
	var user User
	DB.Unscoped().Where("domain = ?", domain).First(&user)
	return user
}
