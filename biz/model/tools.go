package model

import (
	"DouYin-Proj/biz/model/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
	"time"
)

func Connect2sql() *gorm.DB {
	db, err := gorm.Open(
		mysql.Open("root:XXXXXX@tcp(127.0.0.1:3306)/douyin?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}
	return db
}

func Findusername(db *gorm.DB, username string) bool {
	var userinfo user.Userinfo
	db.Where("username = ?", username).Find(&userinfo)
	if userinfo.Username == "" {
		return false
	} else {
		return true
	}
}

func Getuser(db *gorm.DB, username string) user.Userinfo {
	var userinfo user.Userinfo
	db.Where("username = ?", username).Find(&userinfo)
	return userinfo
}

func Encodetoken(userid string, username string) string {
	keyinfo := []byte("3.1415926+0.618+qweasd")
	temp := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":     userid,
		"username":   username,
		"expiretime": time.Now().Unix() + 3600*24,
		"issuer":     "daniel",
	})
	token, err := temp.SignedString(keyinfo)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return token
}

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

func Checktoken(c *app.RequestContext) (bool, *user.Userinfo) {
	token := c.Query("token")
	p := Decodetoken(token)
	if p[0] == "" {
		return false, nil
	}
	userid, _ := strconv.Atoi(p[0])
	var userinfo user.Userinfo
	Connect2sql().Where("ID = ? AND username = ?", userid, p[1]).Find(&userinfo)
	if userinfo.Username != "" {
		return true, &userinfo
	} else {
		return false, nil
	}
}
