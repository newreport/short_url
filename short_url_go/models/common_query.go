package models

import (
	"fmt"
	"reflect"
	"short_url_go/common"
)

// 查找数据库的一条记录
func FirstOrDefault[T SqlModel](query interface{}, args ...interface{}) (one *T) {
	common.DB.Where(query, args).First(&one)
	fmt.Println("fisrt success")
	return
}

// 常用sqlwhere
func Where[T SqlModel](query ...interface{}) (list *[]T) {
	if len(query) > 0 {
		fmt.Println(reflect.TypeOf(query[0]))
	}
	if len(query) == 0 { //查询所有
		common.DB.Find(&list)
	} else if len(query) == 1 { //查找
		switch query[0].(type) {
		case *T, []int64, map[string]interface{}:
			common.DB.Where(query[0]).Find(&list)
		}
	} else if len(query) == 2 { //查找，排序/选择
		// typeQ1:=reflect.Type()
		switch query[0].(type) {
		case *T, []int64, map[string]interface{}:
			switch query[1].(type) {
			case string:
				common.DB.Where(query[0]).Order(query[1]).Find(&list)
			case []string:
				common.DB.Where(query[0]).Select(query[1]).Find(&list)
			}
		}
	} else if len(query) == 3 { //查找，排序,选择/limit
		switch query[0].(type) {
		case *T, []int64, map[string]interface{}:
			switch query[1].(type) {
			case string:
				switch query[2].(type) {
				case []string:
					common.DB.Where(query[0]).Order(query[1]).Select(query[2]).Find(&list)
				case int:
					common.DB.Where(query[0]).Order(query[1]).Limit(int(query[2].(int))).Find(&list)
				}
			}
		}
	}
	return
}
