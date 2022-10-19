package common

import (
	"github.com/beego/beego/v2/core/config"
	"gorm.io/gorm"
)

var DB *gorm.DB

var INIconf config.Configer
