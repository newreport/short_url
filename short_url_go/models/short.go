package models

import (
	"encoding/binary"
	"fmt"
	"short_url_go/common"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Short struct {
	Sid          string    `gorm:"primaryKey,size:50;"`  //主键uuid
	SourceUrl    string    `gorm:"not null"`             //需要跳转的url
	SourceUrlMD5 string    `gorm:"not null"`             //需要跳转url的MD5
	TargetUrl    string    `gorm:"not null;uniqueIndex"` //目标URL
	Remarks      string    //备注
	FkUser       uint      `gorm:"not null"` //外键关联用户
	FKShortGroup uint      `gorm:"not null"` //外键关联分组
	ExpireAt     time.Time //过期时间
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

// 添加一条URL短链接
func AddOneUrlDefault(url string, userId int) {

}

// 添加一条指定长度的短链接
func AddOneUrlAssignLength(url string, userId int, lengthNum int) {

}

// 添加一条自定义长度的短链接
func AddOneUrl(sourceUrl string, targetUrl string) {

}

// URL种子，即浏览器支持的非转义字符，这里只取了64个
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.~!*
const URLSTRS = "LMndefNq3~ZaUVWvw4sQRABCY56rHz0DEFJ127KxyX89IbcPhijklmGS-TgtOopu"

// @title			AddShortUrlUserdefine
// @description		生成自定义的短链接,直接插入数据库
// @auth			sfhj
// @date			2022-10-18
// @param     		data        		map[string]string		"需要生成短链接的source和target集合"
// @param     		userId        		uint					"关联用户主键"
// @param     		shortGroupId    	uint					"分组主键"
// @return			alreadyResult		map[string]string		"已经存在于数据库中的链接集合"
// @return			repeatResult		map[string]string		"目标target重复的集合"
func AddShortUrlUserdefine(data map[string]string, userId uint, shortGroupId uint) (alreadyResult map[string]string, repeatResult map[string]string) {
	keys := common.GetMapAllKeys(data)
	vlues := common.GetMapAllValues(data)
	var keysMD5 []string
	for i := 0; i < len(keys); i++ {
		keysMD5 = append(keysMD5, common.MD5(keys[i]))
	}
	shortsAlready := Where[Short](map[string]interface{}{"source_url_md5": keysMD5}, []string{"source_url", "target_url"})
	for i := 0; i < len(shortsAlready); i++ {
		alreadyResult[shortsAlready[i].SourceUrl] = shortsAlready[i].TargetUrl
	}
	shortsRepeat := Where[Short](map[string]interface{}{"target_url": vlues}, []string{"source_url", "target_url"})
	for i := 0; i < len(shortsRepeat); i++ {
		repeatResult[shortsRepeat[i].SourceUrl] = shortsRepeat[i].TargetUrl
	}
	var installShorts []Short
	for k, v := range data {
		if _, ok := alreadyResult[k]; !ok {
			if _, ok := repeatResult[v]; !ok {
				var err error
				one := Short{Sid: uuid.Must(uuid.NewV4(), err).String(), SourceUrl: k, TargetUrl: GenerateUrlDefault(k), Remarks: "备注", SourceUrlMD5: common.MD5(k), FkUser: userId, FKShortGroup: shortGroupId}
				if err != nil {
					panic("failed to connect database")
				}
				installShorts = append(installShorts, one)
			}
		}
	}
	common.DB.Create(&installShorts)
	return
}

func AddShortUrl() {

}

func AddShortUrls() {

}

// @title			GenerateUrlDefault
// @description		生成单个短链接url，默认6位
// @auth			sfhj
// @date			2022-10-17
// @param     		urls        string		"需要生成短链接的url"
// @return			result		string		"生成后的短链接(查找到的)"
func GenerateUrlDefault(urls string) string {
	return generateUrl(urls, 6)
}

func GenerateUrl(url string, length int) string {
	return generateUrl(url, length)
}

// @title 生成单个url
func generateUrl(url string, length int) (result string) {

	md5Url := common.MD5(url)
	var count int64
	common.DB.Model(&Short{}).Where("source_url_md5 = ? ", md5Url).Count(&count)
	if count > 0 { //存在记录，直接使用
		var one Short
		common.DB.Where("source_url_md5 = ?", md5Url).First(&one)
		result = one.TargetUrl
		return
	}
	//不存在，开始生成
	md5Arr := []string{md5Url[0:8], md5Url[8:16], md5Url[16:24], md5Url[24:32]}
	for i := 0; i < len(md5Arr); i++ {
		num, _ := strconv.ParseUint(md5Arr[i], 16, 32)
		fmt.Print("num:", num)
		by := UInt32ToBytes(uint32(num))
		//将数据密度提升至16个字节，即短url支持4~16长度，有(4*5*6*....*15*16)^64种可能性
		by = append(by, by[0]&by[1])
		by = append(by, by[1]|by[2])
		by = append(by, by[2]^by[3])
		by = append(by, by[3]&^by[0])

		by = append(by, by[0]&by[2])
		by = append(by, by[1]|by[3])
		by = append(by, by[2]^by[0])
		by = append(by, by[3]&^by[1])

		by = append(by, by[0]&by[3])
		by = append(by, by[1]|by[0])
		by = append(by, by[2]^by[1])
		by = append(by, by[3]&^by[2])

		fmt.Println("   ;smbit：", by)
		for j := 0; j < length; j++ {
			urlNum := by[j] % 64
			result += string(URLSTRS[urlNum])
		}
		common.DB.Model(&Short{}).Where("target_url = ? ", result).Count(&count)
		if count == 0 {
			return //不重复，则直接结束循环
		}
	}
	return
}

// @title			GenerateUrlsDefault
// @description		生成多个短链接url，默认6位
// @auth			sfhj
// @date			2022-10-18
// @param     		urls        []string				"需要生成短链接的url"
// @return			result		map[string]string		"生成后的键值对集合"
func GenerateUrlsDefault(urls []string) map[string]string {
	return generateUrls(urls, 6)
}

func GenerateUrls(urls []string, length int) map[string]string {
	return generateUrls(urls, length)
}

// 生成多个url
func generateUrls(urls []string, length int) (result map[string]string) {
	var md5Urls []string
	for i := 0; i < len(urls); i++ {
		md5Urls = append(md5Urls, common.MD5(urls[i]))
	}

	var count int64
	common.DB.Model(&Short{}).Where("source_url_md5 IN ? ", md5Urls).Count(&count)
	if count > 0 { //存在记录，直接使用
		var alreadyShort []Short
		common.DB.Where("source_url_md5 IN ?", md5Urls).Select([]string{"source_url", "target_url"}).Find(&alreadyShort)
		for i := 0; i < len(alreadyShort); i++ {
			result[alreadyShort[i].SourceUrl] = alreadyShort[i].TargetUrl
		}
	}

	for i := 0; i < len(urls); i++ {
		md5Url := common.MD5(urls[i])
		if _, ok := result[urls[i]]; ok { //存在，直接下一次循环
			continue
		}
		//不存在，开始生成
		md5Arr := []string{md5Url[0:8], md5Url[8:16], md5Url[16:24], md5Url[24:32]}
		for j := 0; j < len(md5Arr); j++ {
			// fmt.Print("value:", md5Arr[j])
			// fmt.Print("   ;转int:")
			num, _ := strconv.ParseUint(md5Arr[j], 16, 32)
			// fmt.Print("num:", num)
			by := UInt32ToBytes(uint32(num))
			//将数据密度提升至16个字节，即短url支持4~16长度，有(4*5*6*....*15*16)^64种可能性
			by = append(by, by[0]&by[1])
			by = append(by, by[1]|by[2])
			by = append(by, by[2]^by[3])
			by = append(by, by[3]&^by[0])

			by = append(by, by[0]&by[2])
			by = append(by, by[1]|by[3])
			by = append(by, by[2]^by[0])
			by = append(by, by[3]&^by[1])

			by = append(by, by[0]&by[3])
			by = append(by, by[1]|by[0])
			by = append(by, by[2]^by[1])
			by = append(by, by[3]&^by[2])

			// fmt.Println("   ;smbit：", by)
			for k := 0; k < length; k++ {
				urlNum := by[k] % 64
				result[urls[i]] += string(URLSTRS[urlNum])
			}
			common.DB.Model(&Short{}).Where("target_url = ? ", result).Count(&count)
			if count == 0 {
				break //不重复，则直接结束循环
			}
		}
	}
	return
}

// 无符号int转byte数组
func UInt32ToBytes(i uint32) []byte {
	var buf = make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, i)
	return buf
}
