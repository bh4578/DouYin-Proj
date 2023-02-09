package favorate_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	//点赞操作
	r.POST("/favorite/action/", router.ChecktokenMw(), handler.FavAction)
	//喜欢列表
	r.GET("/favorite/list/", router.ChecktokenMw(), handler.FavList)

}
