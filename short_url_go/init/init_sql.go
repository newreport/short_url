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
		db, err := gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		//迁移数据表(初始化)
		pwdUUID, err := iniconf.String("uuid")
		if err != nil {
			panic(err)
		}
		fmt.Println("uuid:" + pwdUUID)
		u5 := uuid.Must(uuid.FromString(pwdUUID))
		db.AutoMigrate(&models.User{Name: "admin", NickName: "admin", Passwd: uuid.NewV5(u5, pwdUUID).String(), Role: 1, Remarks: "默认管理员"}, &models.Short{SourceUrl: "baidu.com", Remarks: "备注", FkUser: 0})

		//https://www.bookstack.cn/read/beego-2.0-zh/quickstart.md	
		//初始化url
		// urlOne := models.Short{Si}

	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
