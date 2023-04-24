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
		mysql.Open("root:XXXX@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"),
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
	// keyinfo是JWT签发和验证时的密钥（Key），是由一串随机字符串组成。
	keyinfo := []byte("3.1415926+0.618+qweasd")
	// 用于创建Token对象的函数，method表示签名算法，claims表示添加到Token中的声明
	// 返回一个Token对象
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
	// 将 JWT Token 序列化成字符串并对其进行签名，生成字符串token
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
	// jwt.Parse()函数是用来解析JWT字符串并验证签名的函数。
	// 第一个参数时要解析的JWT字符串，第二个参数是一个回调函数，用于解析所需密钥或是证书。
	// 回调函数返回值是验证JWT签名的密钥或证书，JWT库会用返回的密钥对JWT进行解密，返回值为nil时JWT拒接解析并返回错误信息。
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
	if id1 == id2 {
		return true
	}
	var relaion Userrelation
	result := Connect2sql().Where("userid = ? AND targetid = ? AND valid = true", id1, id2).Find(&relaion)
	if result.RowsAffected < 1 {
		return false
	} else {
		return true
	}

}

func Isfavorite(id1 uint, id2 uint) bool {
	var relaion Favoriteinfo
	result := Connect2sql().Where("userid = ? AND videoid = ? AND valid = true", id1, id2).Find(&relaion)
	if result.RowsAffected < 1 {
		return false
	} else {
		return relaion.Valid
	}
}
func Getfovnum(videoid string) int64 {
	var count int64
	Connect2sql().Table("favoriteinfos").Where("videoid = ? AND valid = ?", videoid, true).Count(&count)
	return count
}
func Getfollownum(userid uint) int64 {
	var count int64
	Connect2sql().Table("Userrelations").Where("userid = ? AND valid = ?", userid, true).Count(&count)
	return count
}

func Getfollowernum(targetid uint) int64 {
	var count int64
	Connect2sql().Table("Userrelations").Where("targetid = ? AND valid = ?", targetid, true).Count(&count)
	return count
}
