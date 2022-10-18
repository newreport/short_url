package init

import (
	"fmt"
	"os"
	"short_url_go/common"

	"short_url_go/models"

	"github.com/beego/beego/v2/core/config"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	iniconf, err := config.NewConfig("ini", "conf/secret.conf")
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
		common.DB, err = gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		//迁移数据表(初始化)
		pwdUUID, err := iniconf.String("UUID::UserPwd")
		if err != nil {
			panic(err)
		}
		fmt.Println("uuid:" + pwdUUID)
		//uuid v5加密
		u5 := uuid.Must(uuid.FromString(pwdUUID))
		common.DB.AutoMigrate(&models.User{}, &models.Short{}, &models.ShortGroup{})
		userAdmin := models.User{Name: "admin", NickName: "admin", Passwd: uuid.NewV5(u5, "admin").String(), Role: 1, DefaultUrlLength: 9, Remarks: "默认管理员"}
		common.DB.Create(&userAdmin)
		userUser1 := models.User{Name: "user", NickName: "user", Passwd: uuid.NewV5(u5, "user").String(), DefaultUrlLength: 9, Role: 1, Remarks: "用户"}
		common.DB.Create(&userUser1)
		shortGroup := models.ShortGroup{Name: "默认组", FkUser: userAdmin.ID}
		common.DB.Create(&shortGroup)
		short := []models.Short{{Sid: uuid.Must(uuid.NewV4(), err).String(), SourceUrl: "baidu.com", TargetUrl: models.GenerateUrlDefault("baidu.com"), Remarks: "备注", SourceUrlMD5: common.MD5("baidu.com"), FkUser: userAdmin.ID, FKShortGroup: shortGroup.ID},
			{Sid: uuid.Must(uuid.NewV4(), err).String(), SourceUrl: "google.com", TargetUrl: models.GenerateUrlDefault("google.com"), Remarks: "备注", SourceUrlMD5: common.MD5("google.com"), FkUser: userAdmin.ID, FKShortGroup: shortGroup.ID}}
		common.DB.Create(&short)

		//https://www.bookstack.cn/read/beego-2.0-zh/quickstart.md
		//初始化url
		// urlOne := models.Short{Si}

	} else {
		common.DB, err = gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}
	models.Test()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
