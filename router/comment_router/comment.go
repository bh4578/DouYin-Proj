package comment_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	//发表评论操作
	r.POST("comment/action/", router.ChecktokenMw(), handler.Publishcomment)
	r.GET("comment/list/", router.ChecktokenMw(), handler.Commentlist)

}
