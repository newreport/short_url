package models

import (
	"errors"
	"short_url_go/utils"
	"strconv"
	"time"

	linq "github.com/ahmetb/go-linq/v3"
	"github.com/beego/beego/logs"
	"github.com/samber/lo"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Short struct {
	ID           string         `json:"id" gorm:"primaryKey,size:50;"` //主键uuid
	SourceURL    string         `json:"sourceURL" gorm:"not null"`     //需要跳转的url
	SourceUrlMD5 string         `json:"sourceMD5" gorm:"not null"`     //需要跳转url的MD5
	TargetURL    string         `json:"targetURL" gorm:"not null"`     //目标URL
	FKUser       uint           `json:"fkUser" gorm:"not null"`        //外键关联用户
	ShortGroup   string         `json:"shortGroup" gorm:"not null"`    //外键关联分组
	IsEnable     bool           `json:"isEnable" gorm:"not null"`      //是否启用
	ExpireAt     time.Time      `json:"exp"`                           //过期时间
	CreatedAt    time.Time      `json:"crt"`
	UpdatedAt    time.Time      `json:"upt"`
	DeletedAt    gorm.DeletedAt `json:"det" gorm:"index"`
	Remarks      string         `json:"remarks"` //备注
}

type AddEditShort struct {
	SourceURL  string    `json:"sourceURL"`  //需要跳转的url
	Automatic  bool      `json:"automactic"` //是否自动生成
	Length     int       `json:"length"`     //自动生成的长度
	TargetURL  string    `json:"targetURL"`  //目标URL
	ShortGroup string    `json:"shortGroup"` //外键关联分组
	IsEnable   bool      `json:"isEnable"`   //是否启用
	ExpireAt   time.Time `json:"exp"`        //过期时间
	Remarks    string    `json:"remarks"`    //备注
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

// @Title			CreateShort
// @Description		生成单个短链接url
// @Auth			sfhj
// @Date			2022-10-23
// @Param     		short		models.Short	"需要生成短链接的url"
// @Param     		length		int				"短链接长度"
// @Return			result		bool			"操作是否成功"
func CreateShort(short Short, length int) error {
	var err error
	short.TargetURL = generateUrl(short.SourceURL, short.FKUser, length)
	var existShort Short
	if DB.Unscoped().Where("target_url = ? AND fk_user = ? ", short.TargetURL, short.FKUser).First(&existShort).RowsAffected > 0 {
		return errors.New("該用戶已存在該短链接")
	}
	short.ID = uuid.Must(uuid.NewV4(), err).String()
	short.SourceUrlMD5 = utils.MD5(short.SourceURL)
	result := DB.Create(&short)
	if result.RowsAffected == 0 {
		return errors.New("數據庫錯誤，创建失败")
	}
	return nil
}

// @Title			CreateShort
// @Description		生成单个短链接url
// @Auth			sfhj
// @Date			2022-10-23
// @Param     		short		models.Short	"需要生成短链接的url"
// @Param     		length		int				"短链接长度"
// @Return			result		bool			"操作是否成功"
func CreateShorts(urls []string, fkUser uint, length int) (result map[string]string) {
	return generateUrls(urls, fkUser, length)
}

// @Title			CreateShortCustom
// @Description		生成单个短链接url，自定义target url
// @Auth			sfhj
// @Date			2022-10-23
// @Param     		short		models.Short	"需要生成短链接的url"
// @Return			result		string			"是否成功"
func CreateShortCustom(short Short) error {
	var err error
	short.ID = uuid.Must(uuid.NewV4(), err).String()
	short.SourceUrlMD5 = utils.MD5(short.SourceURL)
	var count int64
	DB.Model(&Short{}).Unscoped().Where("target_url = ? AND fk_user = ? ", short.TargetURL, short.FKUser).Count(&count)
	if count > 0 { //已存在
		return errors.New("該用戶已存在該短链接")
	}
	result := DB.Create(&short)
	if result.RowsAffected == 0 {
		return errors.New("數據庫錯誤，创建失败")
	}
	return nil
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
		v.SourceUrlMD5 = utils.MD5(v.SourceURL)
		targetStrs = append(targetStrs, v.TargetURL)
		sourceUrlMD5s = append(sourceUrlMD5s, v.SourceUrlMD5)
	}
	var count int64
	fkUser := shorts[0].FKUser
	DB.Model(&Short{}).Unscoped().Where("( target_url IN ? OR source_url_md5 IN ? ) AND fk_user = ?", targetStrs, sourceUrlMD5s, fkUser).Count(&count)
	var existSouceShorts []Short
	var existTargetShorts []Short
	var existSouceURLs []string
	var existTargetURLs []string
	if count > 0 { //已存在
		sourceExpress := DB.Unscoped().Where(" source_url_md5 IN ?  AND fk_user = ?", sourceUrlMD5s, shorts[0].FKUser)
		tagetExpress := DB.Unscoped().Where(" target_url IN ? AND fk_user = ?", targetStrs, shorts[0].FKUser)
		sourceExpress.Select("source_url").Find(&existSouceShorts)
		tagetExpress.Select("source_url").Find(&existTargetShorts)

		linq.From(existSouceShorts).SelectT(func(e Short) string {
			return e.SourceURL
		}).ToSlice(&existSouceURLs) //查詢已存在的sourceURL集合
		linq.From(existTargetShorts).SelectT(func(e Short) string {
			return e.SourceURL
		}).ToSlice(&existTargetURLs) //查詢已經存在的targetURL集合

		linq.From(shorts).WhereT(func(s Short) bool {
			return !lo.Contains[string](existSouceURLs, s.SourceURL)
		}).ToSlice(&shorts) //移除已經存在的sourceURL (數據庫不創建)

		linq.From(shorts).WhereT(func(s Short) bool {
			return !lo.Contains[string](existTargetURLs, s.SourceURL)
		}).ToSlice(&shorts) //移除已經存在的targetURL
	}
	DB.Create(&shorts)
	if err != nil {
		panic("failed to add one assign length short url")
	}
	return
}

func UpdateShort(short Short) error {
	var existShort Short
	result := DB.Where("id = ? ", short.ID).First(&existShort)
	if result.RowsAffected > 0 {
		var count int64
		DB.Model(&Short{}).Unscoped().Where("target_url = ? AND fk_user = ? AND target_url != ?", short.TargetURL, short.FKUser, existShort.TargetURL).Count(&count)
		if count > 0 {
			return errors.New("短链接重复")
		}
		existShort = short
		result := DB.Save(&existShort)
		if result.RowsAffected > 0 {
			return nil
		}
		return errors.New("數據庫錯誤，修改失败")
	} else {
		return errors.New("没有查询到该链接")
	}
}

// @Title	RecoverShort
// @Description	從回收站移除
// @Auth	sfhj
// @Date	2022-12-03
// @Param	id	string
func RecoverShort(id string) bool {
	// 条件更新
	result := DB.Model(&Short{}).Unscoped().Where("id = ?", id).Update("deleted_at", nil)
	return result.RowsAffected > 0
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
		if page.Unscoped {
			express = express.Unscoped().Where(" deleted_at IS NULL ")
		}
		express.Order(page.Sort).Select("id,source_url,source_url_md5,target_url,fk_user,short_group,is_enable,created_at,updated_at,expire_at,remarks").Find(&result)
	} else {
		err = errors.New("查詢參數錯誤")
	}
	return
}

func QueryAllByUserID(userID uint) map[string]string {
	var shorts []Short
	DB.Where(" fk_user = ?", userID).Select("source_url", "target_url").Find(&shorts)
	result := make(map[string]string, len(shorts))

	linq.From(shorts).SelectT(func(e Short) map[string]string {
		return map[string]string{e.TargetURL: e.SourceURL}
	}).ToMap(&result) //查詢已存在的sourceURL集合
	return result
}

func QueryShortByID(id string) Short {
	var one Short
	DB.Model(&Short{}).Unscoped().Where("id = ?", id).First(&one)
	return one
}

// @Title DeletedShortUrlByID
// @Description	根據id刪除url
func DeletedShortUrlByID(id string, isUnscoped bool) bool {
	express := DB.Where("id = ?", id)
	if isUnscoped {
		express = express.Unscoped()
	}
	result := express.Delete(&Short{})
	return result.RowsAffected > 0
}

// @Title DeletedShortUrlByIDs
// @Description 根據多個id刪除url
func DeletedShortUrlByIDs(ids []string, isUnscoped bool) bool {
	express := DB.Where(" id IN ?", ids)
	if isUnscoped {
		express = express.Unscoped()
	}
	result := express.Delete(&Short{})
	return result.RowsAffected > 0
}

// @Title			GenerateUrlDefault
// @Description		生成单个短链接url，默认6位
// @Auth			sfhj
// @Date			2022-10-17
// @Param     		urls        string		"需要生成短链接的url"
// @Param     		length       int		"短链接长度"
// @Return			result		string		"生成后的短链接(查找到的)"
func generateUrl(url string, fkUser uint, length int) (result string) {
	logs.Info(url)
	md5Url := utils.MD5(url)
	var count int64
	DB.Model(&Short{}).Unscoped().Where("source_url_md5 = ? AND fk_user = ? ", md5Url, fkUser).Count(&count)
	if count > 0 { //存在记录，直接使用
		var one Short
		DB.Unscoped().Where("source_url_md5 = ?  AND fk_user = ? ", md5Url, fkUser).First(&one)
		result = one.TargetURL
		return
	}
	//不存在，开始生成
	md5Arr := []string{md5Url[0:8], md5Url[8:16], md5Url[16:24], md5Url[24:32]}
	for i := 0; i < len(md5Arr); i++ {
		num, _ := strconv.ParseUint(md5Arr[i], 16, 32)
		// fmt.Print("num:", num)
		by := utils.UInt32ToBytes(uint32(num))
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
		DB.Model(&Short{}).Unscoped().Where("target_url = ? AND fk_user = ? ", result, fkUser).Count(&count)
		if count == 0 {
			return //不重复，则直接结束循环
		}
		result = ""
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
func generateUrls(urls []string, fkUser uint, length int) (result map[string]string) {
	var md5Urls []string
	for i := 0; i < len(urls); i++ {
		md5Urls = append(md5Urls, utils.MD5(urls[i]))
	}
	result = make(map[string]string)
	var count int64
	DB.Model(&Short{}).Unscoped().Where("source_url_md5 IN ?  AND fk_user = ? ", md5Urls, fkUser).Count(&count)
	if count > 0 { //存在记录，直接使用
		var alreadyShort []Short
		DB.Unscoped().Where("source_url_md5 IN ?  AND fk_user = ?", md5Urls, fkUser).Select([]string{"source_url", "target_url"}).Find(&alreadyShort)
		for i := 0; i < len(alreadyShort); i++ {
			result[alreadyShort[i].SourceURL] = alreadyShort[i].TargetURL
		}
	}

	for i := 0; i < len(urls); i++ {
		if _, ok := result[urls[i]]; ok { //存在，直接下一次循环
			continue
		}
		md5Url := utils.MD5(urls[i])
		//不存在，开始生成
		md5Arr := []string{md5Url[0:8], md5Url[8:16], md5Url[16:24], md5Url[24:32]}
		for j := 0; j < len(md5Arr); j++ {
			num, _ := strconv.ParseUint(md5Arr[j], 16, 32)
			by := utils.UInt32ToBytes(uint32(num))
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

			for k := 0; k < length; k++ {
				urlNum := by[k] % 64
				result[urls[i]] += string(URLSTRS[urlNum])
			}
			DB.Model(&Short{}).Unscoped().Where("target_url = ? AND fk_user = ? ", result, fkUser).Count(&count)
			if count == 0 {
				break //不重复，则直接结束循环
			}
		}
	}
	return
}
