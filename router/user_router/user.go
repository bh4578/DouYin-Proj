package user_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

// user相关的路由注册

func Register(r *gin.RouterGroup) {
	// 中间件函数：router.ChecktokenMw可以在路由函数被调用之前或之后执行，从而在请求到达目标处理函数之前或之后完成某些任务，比如：鉴权、日志记录、数据校验等。
	// 路由函数：handler.UserInfo，返回用户信息
	r.GET("/user/", router.ChecktokenMw(), handler.UserInfo)
	// POST函数是用来创建HTTP POST请求路由的
	// 为不同的HTTP POST请求路径指定对应的处理函数，以响应来自客户端的HTTP POST请求。
	// 注册功能：username     password（SHA256加密密码）
	r.POST("/user/register/", router.RegisterMw(), handler.Register)
	// 登录功能：
	r.POST("/user/login/", router.LoginMw(), handler.Login)

}
