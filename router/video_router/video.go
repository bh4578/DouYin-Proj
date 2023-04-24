package video_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	// 浏览器向服务器发起GET请求时，JSON格式返回一些视频内容。
	r.GET("/feed/", router.FeedMw(), handler.Feed)
	// 用户上传视频文件处理
	// 中间件函数：鉴权、提取用户信息
	r.POST("/publish/action/", router.ChecktokenMw(), handler.Publish)
	// 获取指定用户发布的视频列表
	r.GET("/publish/list/", router.ChecktokenMw(), handler.Publishlist)
}
