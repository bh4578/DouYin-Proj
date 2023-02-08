package video_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/feed/", router.FeedMw(), handler.Feed)
	r.POST("/publish/action/", router.ChecktokenMw(), handler.Publish)
	r.GET("/publish/list/", router.ChecktokenMw(), handler.Publishlist)
}
