package user_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

// user相关的路由注册

func Register(r *gin.RouterGroup) {

	r.GET("/user/", router.ChecktokenMw(), handler.UserInfo)
	r.POST("/user/register/", router.RegisterMw(), handler.Register)
	r.POST("/user/login/", router.LoginMw(), handler.Login)

}
