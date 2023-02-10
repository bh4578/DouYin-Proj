package message_router

import (
	"Douyin/handler"
	"Douyin/router"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.RouterGroup) {
	r.POST("/message/action/", router.ChecktokenMw(), handler.Sendmessage)
	r.GET("/message/chat/", router.ChecktokenMw(), handler.Getmessagelist)

}
