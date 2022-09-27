package main

import (
	"database/sql"
	"fmt"
	_ "golang/routers"
	"os"

	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	str, _ := os.Getwd()
	fmt.Println(str)
	InitOpen()

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}

// https://www.jianshu.com/p/c2f72fc0bdb5
func InitOpen() {
	db, err := sql.Open("sqlite3", "./user.db")
	checkErr(err)

	fmt.Println("创建数据表")
	sql_table := `
	CREATE TABLE IF NOT EXISTS "student"(
		"name" VARCHAR(64) NULL,
		"age" VARCHAR(64) NULL,
		"class" VARCHAR(64) NULL
	)`
	db.Exec(sql_table)

	//插入数据
	stmt, err := db.Prepare("INSERT INTO student(name, age, class) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("小明", "12", "六年级一班")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
