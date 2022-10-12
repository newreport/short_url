package models

import (
	"fmt"
	"short_url_go/common"
)

// Sql类型限定
type SqlModel interface {
	User | Short
}

// 基接口
type BaseSqlInterface[T SqlModel] interface {
	FirstOrDefault() *T
}

// 空基struct
type baseSqlStruct struct{}

func (T baseSqlStruct) FirstOrDefault() *User {
	var newU User
	common.DB.First(&newU)
	fmt.Println("fisrt success 继承")
	fmt.Println(newU)
	return &newU
}
