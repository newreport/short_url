package models

import (
	"short_url_go/common"
)

type SqlModel interface {
	User | Short
}

func FirstOrDefault[T SqlModel]() T {
	var model T
	common.DB.First(&model)
	return model
}
