package models

import (
	"fmt"
	"os"
	"short_url_go/common"

	"github.com/beego/beego/v2/core/config"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Page struct {
	PageNum  int    `form:"page_num"`
	PageSize int    `form:"page_size"`
	Keyword  string `form:"keyword"`
	Desc     bool   `form:"desc"`
}

var U5Seed uuid.UUID

func init() {
	var err error
	common.INIconf, err = config.NewConfig("ini", "conf/secret.conf")
	if err != nil {
		checkErr(err)
	}
	pwdUUID, err := common.INIconf.String("UUID::UserPwd")
	if err != nil {
		checkErr(err)
	}
	U5Seed = uuid.Must(uuid.FromString(pwdUUID))

	//1.创建data文件夹，用于存放数据
	_path := "./data"
	existDic, err := common.PathExists(_path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		checkErr(err)
	}
	if !existDic {
		os.Mkdir(_path, os.ModePerm)
	}
	_path += "/main.db"
	if existSqlFile, _ := common.PathExists(_path); !existSqlFile {
		//创建sqlite文件
		DB, err = gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		//https://www.bookstack.cn/read/beego-2.0-zh/quickstart.md
		DB.AutoMigrate(&User{}, &Short{})
		userAdmin := User{Name: "admin", Nickname: "admin", Password: uuid.NewV5(U5Seed, "admin").String(), Role: 1, DefaultURLLength: 6, Remarks: "默认管理员"}
		DB.Create(&userAdmin)
		userUser1 := User{Name: "user", Nickname: "user", Password: uuid.NewV5(U5Seed, "user").String(), DefaultURLLength: 9, Role: 0}
		DB.Create(&userUser1)
	} else {
		DB, err = gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
