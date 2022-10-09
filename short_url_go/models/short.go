package models

type Short struct {
	Sid       string `gorm:"primaryKey,size:50;"`
	SourceUrl string `gorm:"not null"`
	TargetUrl string `gorm:"not null"`
	Remarks   string
	FkUser    uint `gorm:"not null"` //外键
	UrlGroup  string
}

func AddShor(url string, userId int) {

}
