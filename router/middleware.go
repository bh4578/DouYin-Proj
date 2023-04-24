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

// 中间件函数，执行代码由c.Next()分割成路由函数前后执行
// checktoken用于判断token是否正确，通过解码得到的id与用户名在数据库中查询是否存在
// c.Abort()用来阻止后续的处理函数继续执行。在中间件函数中调用该方法，会停止当前请求的处理，并直接返回响应结果。
// c.Set是将名为userinfo的键值对存储在当前请求的上下文中。在后续的中间件函数或处理函数中，可以通过c.Get("userinfo")来获取这个值
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
			c.Abort() // 放弃后续访问，退出函数
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
