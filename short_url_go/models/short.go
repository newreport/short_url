package models

import (
	"short_url_go/common"
	"strconv"
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Short struct {
	Sid          string         `json:"sid" gorm:"primaryKey,size:50;"`        //主键uuid
	SourceURL    string         `json:"sourceUrl" gorm:"not null"`             //需要跳转的url
	SourceUrlMD5 string         `json:"sourceMD5" gorm:"not null"`             //需要跳转url的MD5
	TargetURL    string         `json:"targetURL" gorm:"not null;uniqueIndex"` //目标URL
	Remarks      string         //备注
	FkUser       uint           `json:"fkUser" gorm:"not null"`       //外键关联用户
	FKShortGroup uint           `json:"fkShortGroup" gorm:"not null"` //外键关联分组
	IsEnable     bool           `json:"isEnable" gorm:"not null"`     //是否启用
	ExpireAt     time.Time      `json:"exp"`                          //过期时间
	CreatedAt    time.Time      `json:"crt"`
	UpdatedAt    time.Time      `json:"upt"`
	DeletedAt    gorm.DeletedAt `json:"det" gorm:"index"`
}

// URL种子，即浏览器支持的非转义字符，这里只取了64个
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.~!*
const URLSTRS = "LMndefNq3~ZaUVWvw4sQRABCY56rHz0DEFJ127KxyX89IbcPhijklmGS-TgtOopu"

// @title			CreateShort
// @description		生成单个短链接url
// @auth			sfhj
// @date			2022-10-23
// @param     		short		models.Short	"需要生成短链接的url"
// @param     		length		int				"短链接长度"
// @return			result		string			"生成后的短链接(查找到的)"
func CreateShort(short Short, length int) bool {
	var err error
	short.Sid = uuid.Must(uuid.NewV4(), err).String()
	short.TargetURL = generateUrl(short.TargetURL, length)
	short.SourceUrlMD5 = common.MD5(short.SourceURL)
	result := DB.Create(short)
	if err != nil {
		panic("failed to add one assign length short url")
	}
	return result.RowsAffected > 0
}

// @title			CreateShortCustom
// @description		生成单个短链接url，自定义target url
// @auth			sfhj
// @date			2022-10-23
// @param     		short		models.Short	"需要生成短链接的url"
// @return			result		string			"是否成功"
func CreateShortCustom(short Short) bool {
	var err error
	short.Sid = uuid.Must(uuid.NewV4(), err).String()
	short.SourceUrlMD5 = common.MD5(short.SourceURL)
	var count int64
	DB.Where("target_url = ?", short.TargetURL).Count(&count)
	if count > 0 { //已存在
		return false
	}
	result := DB.Create(short)
	if err != nil {
		panic("failed to add one assign length short url")
	}
	return result.RowsAffected > 0
}

// @title			CreateShortsCustom
// @description		生成多个短链接url，自定义target url
// @auth			sfhj
// @date			2022-10-23
// @param     		short		models.Short	"需要生成短链接的url"
// @return			result		string			"是否成功"
// @return			alreadyResult		map[string]string		"已经存在于数据库中的链接集合"
// @return			repeatResult		map[string]string		"目标target重复的集合"
// func CreateShortsCustom(shorts []Short) (alreadyResult map[string]string, repeatResult map[string]string) {
// 	var err error
// 	var targetStrs []string
// 	var sourceUrlMD5s []string
// 	for _, v := range shorts {
// 		v.Sid = uuid.Must(uuid.NewV4(), err).String()
// 		v.SourceUrlMD5 = common.MD5(v.SourceUrl)
// 		targetStrs = append(targetStrs, v.TargetUrl)
// 		sourceUrlMD5s = append(sourceUrlMD5s, v.SourceUrlMD5)
// 	}
// 	var count int64
// 	DB.Where("target_url IN ? OR source_url_md5 IN ?", targetStrs, sourceUrlMD5s).Count(&count)
// 	if count > 0 { //已存在
// 		return false
// 	}
// 	result := DB.Create(shorts)
// 	if err != nil {
// 		panic("failed to add one assign length short url")
// 	}
// 	return result.RowsAffected > 0
// }

func DeletedShortUrlById(id string) bool {
	result := DB.Delete(&Short{}, id)
	return result.RowsAffected > 0
}

func DeletedShortUrlByIds(ids []string) bool {
	result := DB.Delete(&Short{}, "sid IN ?", ids)
	return result.RowsAffected > 0
}

// @title			GenerateUrlDefault
// @description		生成单个短链接url，默认6位
// @auth			sfhj
// @date			2022-10-17
// @param     		urls        string		"需要生成短链接的url"
// @param     		length       int		"短链接长度"
// @return			result		string		"生成后的短链接(查找到的)"
func generateUrl(url string, length int) (result string) {

	md5Url := common.MD5(url)
	var count int64
	DB.Model(&Short{}).Where("source_url_md5 = ? ", md5Url).Count(&count)
	if count > 0 { //存在记录，直接使用
		var one Short
		DB.Where("source_url_md5 = ?", md5Url).First(&one)
		result = one.TargetURL
		return
	}
	//不存在，开始生成
	md5Arr := []string{md5Url[0:8], md5Url[8:16], md5Url[16:24], md5Url[24:32]}
	for i := 0; i < len(md5Arr); i++ {
		num, _ := strconv.ParseUint(md5Arr[i], 16, 32)
		// fmt.Print("num:", num)
		by := common.UInt32ToBytes(uint32(num))
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
		for j := 0; j < length; j++ {
			urlNum := by[j] % 64
			result += string(URLSTRS[urlNum])
		}
		DB.Model(&Short{}).Where("target_url = ? ", result).Count(&count)
		if count == 0 {
			return //不重复，则直接结束循环
		}
	}
	return
}

// @title			generateUrls
// @description		生成多个短链接url，默认6位
// @auth			sfhj
// @date			2022-10-18
// @param     		urls        []string				"需要生成短链接的url"
// @param     		length        int				"url长度"
// @return			result		map[string]string		"生成后的键值对集合"
func generateUrls(urls []string, length int) (result map[string]string) {
	var md5Urls []string
	for i := 0; i < len(urls); i++ {
		md5Urls = append(md5Urls, common.MD5(urls[i]))
	}

	var count int64
	DB.Model(&Short{}).Where("source_url_md5 IN ? ", md5Urls).Count(&count)
	if count > 0 { //存在记录，直接使用
		var alreadyShort []Short
		DB.Where("source_url_md5 IN ?", md5Urls).Select([]string{"source_url", "target_url"}).Find(&alreadyShort)
		for i := 0; i < len(alreadyShort); i++ {
			result[alreadyShort[i].SourceURL] = alreadyShort[i].TargetURL
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
			by := common.UInt32ToBytes(uint32(num))
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
			DB.Model(&Short{}).Where("target_url = ? ", result).Count(&count)
			if count == 0 {
				break //不重复，则直接结束循环
			}
		}
	}
	return
}
