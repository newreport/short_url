package init

import (
	"fmt"
	"os"
	"short_url_go/common"

	"short_url_go/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	fmt.Println("-------------**************************INITSQL**************************-------------")
	// httpport := beego.AppConfig.String("httpport")
	// fmt.Println("port:" + httpport)

	// fmt.Println("port:" + beego.AppConfig.String("uuid"))
	//1.创建data文件夹，用于存放数据
	_path := "./data"
	existDic, err := common.PathExists(_path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
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

		//创建数据表
		db.AutoMigrate(&models.User{Name: "admin", NickName: "admin", Passwd: common.MD5("admin"), Role: 1, Remarks: "默认管理员"}, &models.Short{})

		// u5 := uuid.Must(uuid.FromString("d4388a9b-594a-7a05-f45b-449b8995819a"))
		//初始化url
		// urlOne := models.Short{Si SourceUrl: "baidu.com", Remarks: "备注", FkUser: 0}

	}
}

func NewConfig(s1, s2 string) {
	panic("unimplemented")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
