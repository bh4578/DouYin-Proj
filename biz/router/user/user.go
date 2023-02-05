package user

import (
	"DouYin-Proj/biz/handler/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// user相关的路由注册

func Register(r *server.Hertz) {

	root := r.Group("/douyin", rootMw()...)
	root.GET("/user/", append(_infoMw(), user.UserInfo)...)
	root.POST("/user/register/", append(_registerMw(), user.Register)...)
	root.POST("/user/login/", append(_loginMw(), user.Login)...)

}
