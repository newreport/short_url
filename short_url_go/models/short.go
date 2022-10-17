package models

import (
	"encoding/binary"
	"fmt"
	"short_url_go/common"
	"strconv"
	"time"

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

const URLSTRS = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-~"

func GenerateUrlDefault(str string) string {
	return generateUrl(str, 6)
}

func generateUrl(url string, length int) (result string) {
	md5Url := common.MD5(url)
	var count int64
	common.DB.Model(&Short{}).Where("source_url_md5 = ? ", md5Url).Count(&count)
	if count > 0 { //存在记录，直接使用
		var one Short
		common.DB.First(&one, "source_url_md5 = ?", md5Url)
		result = one.TargetUrl
		fmt.Println("target:", result)
		return
	}
	md5Arr := []string{md5Url[0:8], md5Url[8:16], md5Url[16:24], md5Url[24:32]}
	for i := 0; i < len(md5Arr); i++ {
		fmt.Print("value:", md5Arr[i])
		fmt.Print("   ;转int:")
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
		result = ""
		for j := 0; j < length; j++ {
			urlNum := by[j] % 64
			result += string(URLSTRS[urlNum])
		}
		common.DB.Model(&Short{}).Where("target_url = ? ", result).Count(&count)
		if count == 0 {
			return
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
