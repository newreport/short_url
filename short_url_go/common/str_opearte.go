package common

import (
	"crypto/md5"
	"fmt"
	"strings"
)

// MD5加密
func MD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	return strings.ToUpper(md5str)
}
