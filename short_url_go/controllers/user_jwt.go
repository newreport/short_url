package controllers

import (
	"short_url_go/models"
	"short_url_go/utils"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/golang-jwt/jwt/v4"
)

// @Title analysisAccountClaims
// @Description 從 http head 中解析 AccessTokenKey
// @Param b beego.Controller	http 上下文
// @Return result  controllers.AccountClaims
func (b *BaseController) analysisAccountClaims() (result AccountClaims) {
	tokenString := b.Ctx.Input.Header("Authorization")
	if len(tokenString) > 0 {
		tokenString = strings.Split(tokenString, " ")[1]
		token, _ := jwt.ParseWithClaims(tokenString, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
			key, _ := utils.INIconf.String("JWT::AccessTokenKey")
			return []byte(key), nil
		})
		logs.Info("token ok:")
		if claims, ok := token.Claims.(*AccountClaims); ok && token.Valid {
			result = *claims
		}
	}
	return
}

func AuthenticationJWT(tokenString string) bool {
	token, _ := jwt.ParseWithClaims(tokenString, &AccountClaims{}, func(token *jwt.Token) (interface{}, error) {
		key, _ := utils.INIconf.String("JWT::AccessTokenKey")
		return []byte(key), nil
	})
	if _, ok := token.Claims.(*AccountClaims); ok && token.Valid {
		return true
	}
	return false
}

// @Title generateAccountJWT
// @Auth sfhj
// @Date 2022-10-26
// @Param user models.User 用户模型
// @Return accountJWT string 请求用的JWT字符串
func generateAccountJWT(user models.User) string {
	claims := AccountClaims{
		user.ID,
		user.Name,
		user.Nickname,
		user.Role,
		user.DefaultURLLength,
		user.I18n,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(10 * time.Minute)), //10分组刷新一次
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "newreport",
			Subject:   "somebody",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, _ := utils.INIconf.String("JWT::AccessTokenKey")
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

type AccountClaims struct {
	ID               uint   `json:"id"`
	Name             string `json:"name"`
	Nickname         string `json:"nickname"`
	Role             int8   `json:"role"`
	DefaultURLLength uint8  `json:"urlLength"`
	I18n             string `json:"i18n"`
	jwt.RegisteredClaims
}

// @Title generateRefreshJWT
// @Auth sfhj
// @Date 2022-10-26
// @Param id uint 用户id
// @Return refreshJWT string 刷新用的JWT字符串
func generateRefreshJWT(id uint) string {
	claims := RefreshClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * 24 * time.Hour)), //15天刷新一次
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "newreport",
			Subject:   "somebody",
			Audience:  []string{"somebody_else"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	key, _ := utils.INIconf.String("JWT::RefreshTokenKey")
	tokenString, _ := token.SignedString([]byte(key))
	return tokenString
}

type RefreshClaims struct {
	ID uint `json:"id"`
	jwt.RegisteredClaims
}
