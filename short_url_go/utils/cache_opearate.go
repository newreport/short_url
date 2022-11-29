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
	AC, err = cache.NewCache("file", `{"CachePath":"./cache_file","FileSuffix":"cache","DirectoryLevel":"2","EmbedExpiry":"120"}`)
	if err != nil {
		fmt.Println("NewCache failed, err:", err)
	}
	fmt.Println(AC)
	// put方法：往缓存里存数据
	// ac.Put("name", "优选短链接", 120)
}
