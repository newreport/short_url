package utils

import (
	"fmt"

	"github.com/beego/beego/v2/client/cache"
)

var DoaminUser map[string]uint

var AC cache.Cache

func RefreshDoaminUser() {

}

// https://www.cnblogs.com/hei-ma/articles/13847724.html
func init() {
	var err error
	/*
	   初始化缓存对象，返回值：cache类型的接口（该接口有一系列方法）和错误信息
	   参数：
	       CachePath：缓存的文件目录
	       FileSuffix：文件后缀
	       DirectoryLevel：目录层级
	       EmbedExpiry：过期时间，字符串类型
	*/
	AC, err = cache.NewCache("file", `{"CachePath":"./cache_file","FileSuffix":"cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	if err != nil {
		fmt.Println("NewCache failed, err:", err)
	}
	fmt.Println(AC)
	// put方法：往缓存里存数据
	// ac.Put("name", "优选短链接", 120)
}
