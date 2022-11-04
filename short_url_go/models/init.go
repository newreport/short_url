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

func init() {
	var err error
	common.INIconf, err = config.NewConfig("ini", "conf/secret.conf")

	if err != nil {
		panic(err)
	}

	//1.创建data文件夹，用于存放数据
	_path := "./data"
	existDic, err := common.PathExists(_path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		panic(err)
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

		//迁移数据表(初始化)
		pwdUUID, err := common.INIconf.String("UUID::UserPwd")
		if err != nil {
			panic(err)
		}
		//uuid v5加密
		u5 := uuid.Must(uuid.FromString(pwdUUID))
		fmt.Println("u5:", u5)
		DB.AutoMigrate(&User{}, &Short{}, &ShortGroup{})
		userAdmin := User{Name: "admin", Nickname: "admin", Password: uuid.NewV5(u5, "admin").String(), Role: 1, DefaultURLLength: 9, Remarks: "默认管理员"}
		DB.Create(&userAdmin)
		fmt.Print("admin:")
		fmt.Println(userAdmin.Password)
		userUser1 := User{Name: "user", Nickname: "user", Password: uuid.NewV5(u5, "user").String(), DefaultURLLength: 9, Role: 1, Remarks: "用户"}
		DB.Create(&userUser1)
		//https://www.bookstack.cn/read/beego-2.0-zh/quickstart.md
		//初始化url
		// urlOne := Short{Si}

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
