package main

import (
	"database/sql"
	"fmt"
	common "goloang/common"
	_ "goloang/routers"
	"os"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {

	{ //1.创建data文件夹，用于存放数据
		_dir := "./data"
		exist, err := common.PathExists(_dir)
		if err != nil {
			fmt.Printf("get dir error![%v]\n", err)
			return
		}
		if !exist {
			os.Mkdir(_dir, os.ModePerm)
		}
	}

	{ //2.初始化sqlite数据库https://www.jianshu.com/p/c2f72fc0bdb5

		//创建sqlite文件
		db, err := sql.Open("sqlite3", "./data/main.db")
		checkErr(err)

		//创建user表
		sql_table := `
		create table user
		(
			uid     INTEGER not null
				constraint user_pk
					primary key autoincrement,
			name    TEXT    not null,
			passwd  TEXT    not null,
			role    INTEGER not null,
			remarks TEXT
		);`
		db.Exec(sql_table)

		//插user默认数据
		stmt, err := db.Prepare(`insert into user (uid, name, passwd, role, remarks) values (0,"admin","admin",1,"默认管理员");`)
		checkErr(err)

		res, err := stmt.Exec()
		checkErr(err)
		

		id, err := res.LastInsertId()
		checkErr(err)
		fmt.Println(id)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
