package models

import (
	"fmt"
	"os"
	"short_url_go/utils"
	"strings"

	"github.com/beego/beego/v2/core/config"
	uuid "github.com/satori/go.uuid"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Page struct {
	Offset   int    `json:"offset"`   //偏移量
	Lmit     int    `json:"limit"`    //指定返回记录的数量
	Sort     string `json:"sort"`     //排序 RHS 例:id+,create_at- ——> id asc,create_at desc
	Unscoped bool   `json:"unscoped"` //回收站
}

var U5Seed uuid.UUID

// URL种子，即浏览器支持的非转义字符，这里只取了64个
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.~!*
var URLSTRS = "LMndeJ127KxyX89IbVWvw4sQRABCY56rHfNq3~ZaUz0DEFcPhijklmGS-TgtOopu"

func init() {
	var err error
	utils.INIconf, err = config.NewConfig("ini", "../data/go/data/secret.conf")
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
	pwdUUID, err := utils.INIconf.String("UUID::UserPwd")
	if err != nil {
		panic(err)
	}
	U5Seed = uuid.Must(uuid.FromString(pwdUUID))

	//1.创建data文件夹，用于存放数据
	_path := "./data"
	existDic, err := utils.PathExists(_path)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		panic(err)
	}
	if !existDic {
		os.Mkdir(_path, os.ModePerm)
	}

	//2.创建sqlite数据库
	_path += "/main.db"
	if existSqlFile, _ := utils.PathExists(_path); !existSqlFile {
		DB, err = gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		//https://www.bookstack.cn/read/beego-2.0-zh/quickstart.md
		DB.AutoMigrate(&User{}, &Short{})
		userAdmin := User{Name: "admin", Nickname: "admin", Password: uuid.NewV5(U5Seed, "admin").String(), Role: 1, DefaultURLLength: 6, Remarks: "默认管理员"}
		DB.Create(&userAdmin)
	} else {
		DB, err = gorm.Open(sqlite.Open(_path), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
	}

	//3.引用URL种子
	URLSTRS, err = utils.INIconf.String("URL::Seed")
	if err != nil {
		fmt.Printf("get URLSeed error![%v]\n", err)
		panic(err)
	}
}

// @Title analysisRestfulRHS
// @Description	解析api参数
// @Param	db	dbContent上下文
// @Param field	數據庫字段名
// @Param	queryRule	查詢條件
// @Return	bool	值是否正確，即前端隻的解析是否正確
func analysisRestfulRHS(db *gorm.DB, field string, queryRule string) bool {
	//RHS式查詢，將查詢條件開放給前端，自由度更大,權限通過token在controller控制，不通過查詢參數
	//SQL注入：https://gorm.io/zh_CN/docs/security.html#%E5%86%85%E8%81%94%E6%9D%A1%E4%BB%B6
	if len(queryRule) > 0 {
		rhsBrackets := strings.Split(queryRule, ",")
		for i := 0; i < len(rhsBrackets); i++ {
			if len(rhsBrackets[i]) > 2 {
				lhs := rhsBrackets[i][:2]
				param := rhsBrackets[i][2:]
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
					db = db.Where(field+" LIKE ? ", param+"%")
				case "af": //like 模糊查詢結尾	*效率低，索引失效
					db = db.Where(field+" LIKE ? ", "%"+param)
				case "lk": //like 模糊查詢	*效率低，索引失效
					db = db.Where(field+" LIKE ? ", "%"+param+"%")
				default: //前端傳遞查詢參數錯誤
					return false
				}
			}
		}
	}
	return true
}
