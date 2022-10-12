package models

import (
	"fmt"
	"short_url_go/common"
)

func FirstOrDefault[T SqlModel]() User {
	var newU User
	common.DB.First(&newU)
	fmt.Printf("fisrt success")
	return newU
}
