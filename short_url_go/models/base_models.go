package models

import (
	"fmt"
	"short_url_go/common"
)

type baseSqlMethod[T SqlModel] interface {
	FirstOrDefault() *T
}

type SqlModel interface {
	User | Short
}

type BaseStruct struct{}

func FirstOrDefault[T SqlModel]() User {
	var newU User
	common.DB.First(&newU)
	fmt.Printf("fisrt success")
	return newU
}

func (T BaseStruct) FirstOrDefault() *User {
	var newU User
	common.DB.First(&newU)
	fmt.Println("fisrt success 继承")
	return &newU
}
