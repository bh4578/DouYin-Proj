// Code generated by hertz generator.

package main

import (
	handler "DouYin-Proj/biz/handler"
	"DouYin-Proj/biz/router/user"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	user.Register(r)
}