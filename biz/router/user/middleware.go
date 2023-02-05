package user

import (
	"DouYin-Proj/biz/handler"
	"DouYin-Proj/biz/handler/user"
	"DouYin-Proj/biz/model"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"net/http"
)

//该包存放中间件的函数

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _infoMw() []app.HandlerFunc {
	// your code
	return []app.HandlerFunc{func(ctx context.Context, c *app.RequestContext) {
		flag, userinfo := model.Checktoken(c)
		if flag {
			c.Set("userinfo", userinfo)
		} else {
			c.JSON(http.StatusOK, user.UserResponse{
				Response: handler.Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
			c.Abort()
		}
	},
	}
}

func _loginMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _registerMw() []app.HandlerFunc {
	// your code...
	return nil
}
