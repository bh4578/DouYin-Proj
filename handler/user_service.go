package handler

import (
	"Douyin/model"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// 返回UserID 与 token 的结构体
type UserLoginResponse struct {
	Response
	UserId uint   `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// 返回用户信息的结构体
type UserResponse struct {
	Response
	User UserLoginInfo `json:"user"`
}

// MD5加密.运行速度快
func MD5(str string) string {
	data := []byte(str) //切片
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}

// sha256加密，运行速度慢，可靠性强
func Sha256(src string) string {
	m := sha256.New()
	m.Write([]byte(src))
	res := hex.EncodeToString(m.Sum(nil))
	return res
}

// 用户注册
func Register(c *gin.Context) {
	username := c.Query("username")
	//需要对密码进行sha256加密(MD5运行速度更快，sha256更安全)，而不是明文存储
	password := Sha256(c.Query("password"))
	//连接数据库
	db := model.Connect2sql()
	//在数据库中查找是否存在该用户名
	if model.Findusername(db, username) {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		//若数据库中不存在该用户名，则创建新用户
		userinfo := model.Userinfo{Username: username, Password: password}
		db.Create(&userinfo)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId:   userinfo.ID,
			//这里token只做了简单处理，可以使用jwt中间件
			Token: model.Encodetoken(strconv.FormatUint(uint64(userinfo.ID), 10), userinfo.Username),
		})
	}
}

// 用户登录
func Login(c *gin.Context) {
	username := c.Query("username")
	//需要对密码进行sha256加密(MD5运行速度更快，sha256更安全)，而不是明文存储
	password := Sha256(c.Query("password"))
	// 该函数就是返回一个db来操作SQL
	db := model.Connect2sql()
	//在数据库中获得该用户
	userinfo := model.Getuser(db, username)
	//如果数据库中存在该用户，则返回需要的信息
	if userinfo.Username != "" {
		if userinfo.Password == password {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userinfo.ID,
				//token只做了简单处理
				Token: model.Encodetoken(strconv.FormatUint(uint64(userinfo.ID), 10), userinfo.Username),
			})
		} else {
			//若数据库中不存在该用户或密码错误，则返回错误信息
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "The user name does not exist or the password is incorrect"},
			})
		}
	} else {
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "The user name does not exist or the password is incorrect"},
		})

	}

}

// 返回用户信息
// Query会解码URL中带有的参数user_id
func UserInfo(c *gin.Context) {
	userid := c.Query("user_id")
	var userinfo model.Userinfo
	//从数据库中根据userid获得该用户，将数据存在userinfo结构体中
	model.Connect2sql().First(&userinfo, userid)
	// 在前边中间件函数c.Set将名为userinfo的键值对存储在当前请求的上下文中。Get获取这个值
	loginuser, _ := c.Get("userinfo")
	userlogininfo := UserLoginInfo{Id: userinfo.ID, Name: userinfo.Username,
		FollowCount: userinfo.FollowCount, FollowerCount: userinfo.FollowerCount,
		IsFollow: model.Isfollow(loginuser.(*model.Userinfo).ID, userinfo.ID),
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0}, User: userlogininfo,
	})

}
