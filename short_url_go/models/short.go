package models

import (
	"errors"
	"fmt"
	"short_url_go/common"
	"strconv"
	"time"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/samber/lo"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Short struct {
	ID           string         `json:"id" gorm:"primaryKey,size:50;"`         //主键uuid
	SourceURL    string         `json:"sourceURL" gorm:"not null"`             //需要跳转的url
	SourceUrlMD5 string         `json:"sourceMD5" gorm:"not null"`             //需要跳转url的MD5
	TargetURL    string         `json:"targetURL" gorm:"not null;uniqueIndex"` //目标URL
	FKUser       uint           `json:"fkUser" gorm:"not null"`                //外键关联用户
	ShortGroup   string         `json:"shortGroup" gorm:"not null"`            //外键关联分组
	IsEnable     bool           `json:"isEnable" gorm:"not null"`              //是否启用
	ExpireAt     time.Time      `json:"exp"`                                   //过期时间
	CreatedAt    time.Time      `json:"crt"`
	UpdatedAt    time.Time      `json:"upt"`
	DeletedAt    gorm.DeletedAt `json:"det" gorm:"index"`
	Remarks      string         `json:"remarks"` //备注
}

type ShortQueryParams struct {
	ID         string `json:"id"`
	SourceURL  string `json:"source_url"`
	TargetURL  string `json:"target_url"`
	FKUser     uint   `json:"-"`
	ShortGroup string `json:"group"`
	IsEnable   string `json:"is_enable"`
	ExpireAt   string `json:"exp"`
	CreatedAt  string `json:"crt"`
	UpdatedAt  string `json:"upt"`
	DeletedAt  string `json:"det"`
}

// URL种子，即浏览器支持的非转义字符，这里只取了64个
// ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_.~!*
const URLSTRS = "LMndefNq3~ZaUVWvw4sQRABCY56rHz0DEFJ127KxyX89IbcPhijklmGS-TgtOopu"

// @Title			CreateShort
// @Description		生成单个短链接url
// @Auth			sfhj
// @Date			2022-10-23
// @Param     		short		models.Short	"需要生成短链接的url"
// @Param     		length		int				"短链接长度"
// @Return			result		bool			"操作是否成功"
func CreateShort(short Short, length int) bool {
	var err error
	short.ID = uuid.Must(uuid.NewV4(), err).String()
	short.TargetURL = generateUrl(short.TargetURL, length)
	short.SourceUrlMD5 = common.MD5(short.SourceURL)
	result := DB.Create(&short)
	if err != nil {
		panic("failed to add one assign length short url")
	}
	return result.RowsAffected > 0
}

// @Title			CreateShortCustom
// @Description		生成单个短链接url，自定义target url
// @Auth			sfhj
// @Date			2022-10-23
// @Param     		short		models.Short	"需要生成短链接的url"
// @Return			result		string			"是否成功"
func CreateShortCustom(short Short) bool {
	var err error
	short.ID = uuid.Must(uuid.NewV4(), err).String()
	short.SourceUrlMD5 = common.MD5(short.SourceURL)
	var count int64
	DB.Model(&Short{}).Where("target_url = ?", short.TargetURL).Count(&count)
	if count > 0 { //已存在
		return false
	}
	result := DB.Create(&short)
	if err != nil {
		panic("failed to add one assign length short url")
	}
	return result.RowsAffected > 0
}

// @Title			CreateShortsCustom
// @Description		生成多个短链接url，自定义target url
// @Auth			sfhj
// @Date			2022-10-23
// @Param     		shorts		models.Short	"需要生成短链接的url集合"
// @Return			alreadyResult		map[string]string		"错误1：sourceURL重复集合"
// @Return			repeatResult		map[string]string		"错误2：targetURL重复集合"
func CreateShortsCustom(shorts []Short) (alreadyResult map[string]string, repeatResult map[string]string) {
	var err error
	var targetStrs []string
	var sourceUrlMD5s []string
	for _, v := range shorts {
		v.ID = uuid.Must(uuid.NewV4(), err).String()
		v.SourceUrlMD5 = common.MD5(v.SourceURL)
		targetStrs = append(targetStrs, v.TargetURL)
		sourceUrlMD5s = append(sourceUrlMD5s, v.SourceUrlMD5)
	}
	var count int64
	var fkUserID uint
	if len(shorts) > 0 && shorts[0].FKUser != 0 {
		fkUserID = shorts[0].FKUser
		DB.Model(&Short{}).Where("( target_url IN ? OR source_url_md5 IN ? ) AND fk_user = ?", targetStrs, sourceUrlMD5s, fkUserID).Count(&count)
	} else {
		DB.Model(&Short{}).Where("target_url IN ? OR source_url_md5 IN ?", targetStrs, sourceUrlMD5s).Count(&count)
	}
	var existSouceShorts []Short
	var existTargetShorts []Short
	var existSouceURLs []string
	var existTargetURLs []string
	even := lo.Filter[int]([]int{1, 2, 3, 4}, func(x int, index int) bool {
		return x%2 == 0
	})
	fmt.Println(even)
	if count > 0 { //已存在
		sourceExpress := DB.Where(" source_url_md5 IN ? ", sourceUrlMD5s)
		tagetExpress := DB.Where(" target_url IN ? ", targetStrs)
		if len(shorts) > 0 && shorts[0].FKUser != 0 { //查詢用戶外鍵關聯
			sourceExpress = sourceExpress.Where("fk_user = ?", shorts[0].FKUser)
			tagetExpress = tagetExpress.Where("fk_user = ?", shorts[0].FKUser)
		}
		sourceExpress.Select("source_url").Find(&existSouceShorts)
		tagetExpress.Select("source_url").Find(&existTargetShorts)
		linq.From(existSouceShorts).SelectT(func(e Short) string { //查詢已存在的sourceURL集合
			return e.SourceURL
		}).ToSlice(&existSouceURLs)
		linq.From(existTargetShorts).SelectT(func(e Short) string { //查詢已經存在的targetURL集合
			return e.SourceURL
		}).ToSlice(&existTargetURLs)
		//移除已經存在的sourceURL (數據庫不創建)
		linq.From(shorts).WhereT(func(s Short) bool {
			return !lo.Contains[string](existSouceURLs, s.SourceURL)
		}).ToSlice(&shorts)
		//移除已經存在的targetURL
		linq.From(shorts).WhereT(func(s Short) bool {
			return !lo.Contains[string](existTargetURLs, s.SourceURL)
		}).ToSlice(&shorts)
	}
	DB.Create(&shorts)
	if err != nil {
		panic("failed to add one assign length short url")
	}
	return
}

//https://www.cnblogs.com/liuhui5599/p/14081524.html

// @Title	QueryShortsPage
// @Auth	sfhj
// @Date	2022-11-14
// @Param	query	models.ShortQueryParams	查询参数
// @Param	page	models.Page	分页查询
// @Return	result
func QueryShortsPage(page Page, fkUser string, sourceURL string, targetURL string, shortGroup string, isEnable string, exp string, crt string, upt string, del string) (result []Short, count int64, err error) {
	express := DB.Model(&Short{})
	if analysisRestfulRHS(express, "fk_user", fkUser) &&
		analysisRestfulRHS(express, "source_url", sourceURL) &&
		analysisRestfulRHS(express, "target_url", targetURL) &&
		analysisRestfulRHS(express, "short_group", shortGroup) &&
		analysisRestfulRHS(express, "is_enable", isEnable) &&
		analysisRestfulRHS(express, "expire_at", exp) &&
		analysisRestfulRHS(express, "created_at", crt) &&
		analysisRestfulRHS(express, "updated_at", upt) &&
		analysisRestfulRHS(express, "deleted_at", del) {
		express.Count(&count)
		express.Order(page.Sort).Select("id,source_url,source_url_md5,target_url,fk_user,short_group,is_enable,expire_at,created_at,remarks").Find(&result)
	} else {
		err = errors.New("查詢參數錯誤")
	}
	return
}

// @Title DeletedShortUrlById
// @Description	根據id刪除url
func DeletedShortUrlById(id string) bool {
	result := DB.Delete(&Short{}, id)
	return result.RowsAffected > 0
}

// @Title DeletedShortUrlByIds
// @Description 根據多個id刪除url
func DeletedShortUrlByIds(ids []string) bool {
	result := DB.Delete(&Short{}, "id IN ?", ids)
	return result.RowsAffected > 0
}

// @Title			GenerateUrlDefault
// @Description		生成单个短链接url，默认6位
// @Auth			sfhj
// @Date			2022-10-17
// @Param     		urls        string		"需要生成短链接的url"
// @Param     		length       int		"短链接长度"
// @Return			result		string		"生成后的短链接(查找到的)"
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
		by = append(by, by[0]&by[1])  //与
		by = append(by, by[1]|by[2])  //非
		by = append(by, by[2]^by[3])  //或
		by = append(by, by[3]&^by[0]) //异或

		by = append(by, by[0]&by[2])
		by = append(by, by[1]|by[3])
		by = append(by, by[2]^by[0])
		by = append(by, by[3]&^by[1])

		by = append(by, by[0]&by[3])
		by = append(by, by[1]|by[0])
		by = append(by, by[2]^by[1])
		by = append(by, by[3]&^by[2])

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

// @Title			generateUrls
// @Description		生成多个短链接url，默认6位
// @Auth			sfhj
// @Date			2022-10-18
// @Param     		urls        []string				"需要生成短链接的url"
// @Param     		length        int				"url长度"
// @Return			result		map[string]string		"生成后的键值对集合"
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
			num, _ := strconv.ParseUint(md5Arr[j], 16, 32)
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
