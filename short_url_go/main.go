package main

import (
	"database/sql"
	"fmt"
	"os"
	_ "short_url_go/routers"

	"short_url_go/common"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

func init() {
	//1.创建data文件夹，用于存放数据
	_dir := "./data"
	exist, err := common.PathExists(_dir)
	if err != nil {
		fmt.Printf("get dir error![%v]\n", err)
		return
	}
	if !exist {
		os.Mkdir(_dir, os.ModePerm) //2.初始化sqlite数据库https://www.jianshu.com/p/c2f72fc0bdb5
		//创建sqlite文件
		db, err := sql.Open("sqlite3", "./data/main.db")
		checkErr(err)

		//创建user表
		sql_table := `
			create table IF NOT EXISTS user
			(
				uid      INTEGER not null
					 constraint user_pk
						 primary key autoincrement,
				name     TEXT    not null,
				nickname TEXT	 not null,
				passwd   TEXT    not null,
				role     INTEGER not null,
				remarks  TEXT
			);`
		db.Exec(sql_table)

		//插user默认数据
		stmt, err := db.Prepare(`insert into user ( name,nickname, passwd, role, remarks) values (?,?,?,?,?);`)
		checkErr(err)
		pwd := common.MD5("admin")
		fmt.Println(pwd)
		stmt.Exec("admin", "admin", pwd, 1, "默认管理员")

		//创建short表
		sql_table = `create table IF NOT EXISTS short
		(
			sid                TEXT    not null
				constraint short_pk
					primary key,
			source_url         TEXT    not null,
			target_url         TEXT    not null,
			create_time        REAL    not null,
			latest_update_time REAL    not null,
			remarks            TEXT,
			fk_user_sid        INTEGER not null,
			url_group          TEXT
		);`
		db.Exec(sql_table)

		db.Close()
	}
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err) //https://zhuanlan.zhihu.com/p/373653492
	}
}
