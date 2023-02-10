package handler

import (
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Sendmessage(c *gin.Context) {

	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	action_type := c.Query("action_type")
	content := c.Query("content")
	loginsuer, _ := c.Get("userinfo")

	if action_type == "1" {
		model.Connect2sql().Create(&model.Message{Content: content, Authorid: loginsuer.(*model.Userinfo).ID, Targetid: uint(to_user_id), Create_date: time.Now()})
	}

	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "已发送信息"})

}

type Responsemessage struct {
	Id         uint   `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}
type ChatResponse struct {
	Response
	MessageList []Responsemessage `json:"message_list"`
}

func Getmessagelist(c *gin.Context) {

	to_user_id := c.Query("to_user_id")
	loginuser, _ := c.Get("userinfo")
	user := loginuser.(*model.Userinfo)
	idlist := []string{to_user_id, strconv.FormatUint(uint64(user.ID), 10)}
	var messagelist []model.Message
	model.Connect2sql().Table("messages").Select("id,authorid,targetid,create_date,content").Where("authorid IN ? OR targetid IN ?", idlist, idlist).Scan(&messagelist)

	reslist := make([]Responsemessage, len(messagelist))

	for index, val := range messagelist {
		reslist[index].Content = val.Content
		reslist[index].Id = val.Id
		reslist[index].CreateTime = val.Create_date.Format("2006-01-02 15:04:05")
	}

	c.JSON(http.StatusOK, ChatResponse{Response: Response{StatusCode: 0}, MessageList: reslist})
}
