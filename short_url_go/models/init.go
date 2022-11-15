package models

import (
	"fmt"
	"os"
	"short_url_go/common"
	"strings"

	"github.com/beego/beego/v2/core/config"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Page struct {
	Offset int    `json:"offset"` //偏移量
	Lmit   int    `json:"limit"`  //指定返回记录的数量
	Sort   string `json:"sort"`   //排序 RHS 例:id+,create_at- ——> id asc,create_at desc
}

var U5Seed uuid.UUID

func init() {
	var err error
	common.INIconf, err = config.NewConfig("ini", "conf/secret.conf")
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
	pwdUUID, err := common.INIconf.String("UUID::UserPwd")
	if err != nil {
		panic(err)
	}
	U5Seed = uuid.Must(uuid.FromString(pwdUUID))

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

// @Title analysisRestfulRHS
// @Description	解析api参数
func analysisRestfulRHS(db *gorm.DB, field string, str string) bool {
	if len(str) > 0 {
		params := strings.Split(str, ",")
		for i := 0; i < len(params); i++ {
			if len(params[i]) > 2 {
				lhs := params[i][:2]
				param := params[i][2:]
				switch lhs {
				case "lt": //less than 小於
					db = db.Where(field+" < ? ", param)
				case "le": //less than or equal to 小於
					db = db.Where(field+" <= ? ", param)
				case "eq": //equal to 等於
					db = db.Where(field+" = ? ", param)
				case "ne": //not equal to 不等於
					db = db.Where(field+" != ? ", param)
				case "ge": //greater than or equal to 大於等於
					db = db.Where(field+" >= ? ", param)
				case "gt": //greater than 大於
					db = db.Where(field+" > ? ", param)
				case "be": //like 模糊查詢頭	*效率低，索引失效
					db = db.Where(field+" LIKE ?% ", param)
				case "af": //like 模糊查詢結尾	*效率低，索引失效
					db = db.Where(field+" LIKE %?", param)
				case "lk": //like 模糊查詢	*效率低，索引失效
					db = db.Where(field+" LIKE %?% ", param)
				default: //前端傳遞查詢參數錯誤
					return false
				}
			}
		}

	}
	return true
}
