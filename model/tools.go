package model

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var _db *gorm.DB

func init() {
	var err error
	_db, err = gorm.Open(
		mysql.Open("root:XXXXXX@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
}
func Connect2sql() *gorm.DB {
	return _db
}

func Findusername(db *gorm.DB, username string) bool {
	var userinfo Userinfo
	db.Where("username = ?", username).Find(&userinfo)
	if userinfo.Username == "" {
		return false
	} else {
		return true
	}
}

func Getuser(db *gorm.DB, username string) Userinfo {
	var userinfo Userinfo
	db.Where("username = ?", username).Find(&userinfo)
	return userinfo
}

// Encodetoken 此函数用于做jwt编码
func Encodetoken(userid string, username string) string {
	keyinfo := []byte("3.1415926+0.618+qweasd")
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":   userid,
		"username": username,
		// exp: jwt的过期时间，这个过期时间必须要大于签发时间
		"exp": time.Now().Unix() + 3600*24,
		// iss: jwt签发者
		"iss": "daniel",
		// nbf: 定义在什么时间之前，该jwt都是不可用的.
		"nbf": time.Now().Unix(),
		// sub: jwt所面向的用户
		// aud: 接收jwt的一方
		// iat: jwt的签发时间
		// jti: jwt的唯一身份标识，主要用来作为一次性token,从而回避重放攻击。
	})
	token, err := temp.SignedString(keyinfo)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return token
}

// Decodetoken 此函数用于做jwt解码，返回解码后得到的用户id与用户名
func Decodetoken(token string) []string {
	keyinfo := []byte("3.1415926+0.618+qweasd")
	parse, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return keyinfo, nil
	})
	if err != nil {
		log.Println(err.Error())
		return []string{"", ""}
	}
	return []string{parse.Claims.(jwt.MapClaims)["userid"].(string), parse.Claims.(jwt.MapClaims)["username"].(string)}

}

// Checktoken 此函数用于判断token是否正确，通过解码得到的id与用户名在数据库中查询是否存在
func Checktoken(c *gin.Context) (bool, *Userinfo) {
	token := c.Query("token")
	if token == "" {
		token = c.PostForm("token")
	}
	p := Decodetoken(token)
	if p[0] == "" {
		return false, nil
	}
	var userinfo Userinfo
	result := Connect2sql().Where("ID = ? AND username = ?", p[0], p[1]).Find(&userinfo)
	if result.RowsAffected > 0 {
		return true, &userinfo
	} else {
		return false, nil
	}
}

// 用于判断id1是否关注id2
func Isfollow(id1 uint, id2 uint) bool {
	var relaion Userrelation
	result := Connect2sql().Where("userid = ? AND targetid = ?", id1, id2).Find(&relaion)
	if result.RowsAffected < 1 {
		return false
	} else {
		return true
	}

}

func Isfavorite(id1 uint, id2 uint) bool {
	var relaion Favoriteinfo
	result := Connect2sql().Where("userid = ? AND videoid = ?", id1, id2).Find(&relaion)
	if result.RowsAffected < 1 {
		return false
	} else {
		return relaion.Valid
	}
}
func Getfovnum(id string) int64 {
	var relaion Favoriteinfo
	res := Connect2sql().Where("videoid = ? AND valid =?", id, true).Find(&relaion)
	return res.RowsAffected
}
