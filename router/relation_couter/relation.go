package relation_couter

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.POST("/relation/action/", router.ChecktokenMw(), handler.Follow)
	r.GET("/relation/follow/list/", router.ChecktokenMw(), handler.Getfollowlist)
	r.GET("/relation/friend/list/", router.ChecktokenMw(), handler.Getfriendlist)
	r.GET("/relation/follower/list/", router.ChecktokenMw(), handler.Getfollowerlist)

}
