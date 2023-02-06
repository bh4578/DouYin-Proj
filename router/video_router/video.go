package video_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.GET("/feed/", router.ChecktokenMw(), handler.Feed)
}
