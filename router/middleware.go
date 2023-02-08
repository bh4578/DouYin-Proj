package router

import (
	"Douyin/handler"
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

//该包存放中间件的函数

func RootMw() gin.HandlerFunc {
	// your code...
	return func(c *gin.Context) {

	}
}

func ChecktokenMw() gin.HandlerFunc {
	// your code
	return func(c *gin.Context) {
		flag, userinfo := model.Checktoken(c)
		if flag {
			c.Set("userinfo", userinfo)
		} else {
			c.JSON(http.StatusOK, handler.UserResponse{
				Response: handler.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			c.Abort()
		}
	}
}

func LoginMw() gin.HandlerFunc {
	// your code...
	return func(c *gin.Context) {

	}
}

func RegisterMw() gin.HandlerFunc {
	// your code...
	return func(c *gin.Context) {

	}
}

func FeedMw() gin.HandlerFunc {
	// your code...
	return func(c *gin.Context) {

	}
}

func PublishMw() gin.HandlerFunc {

	return func(c *gin.Context) {

	}
}
