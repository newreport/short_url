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
		common.DB.AutoMigrate(&models.User{}, &models.Short{})
		user := models.User{Name: "admin", NickName: "admin", Passwd: uuid.NewV5(u5, pwdUUID).String(), Role: 1, Remarks: "默认管理员"}
		common.DB.Create(&user)
		sid := uuid.Must(uuid.NewV4(), err)
		short := models.Short{Sid: sid.String(), SourceUrl: "baidu.com", TargetUrl: "123", Remarks: "备注", FkUser: 0}
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
	model1 := models.FirstOrDefault[models.User]()
	fmt.Println(model1)
	model2 := models.FirstOrDefault[models.Short]()
	fmt.Println(model2)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
