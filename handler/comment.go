package handler

import (
	"Douyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CommentResponse struct {
	ID          uint          `json:"id"`
	User        UserLoginInfo `json:"user" gorm:"foreignKey:Authorid"`
	Content     string        `json:"content"`
	Create_date string        `json:"create_date"`
}

type CommentListResponse struct {
	Response
	CommentList []CommentResponse `json:"comment_list,omitempty"`
}

func Publishcomment(c *gin.Context) {
	video_id, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	action_type := c.Query("action_type")
	comment_text := c.Query("comment_text")
	comment_id := c.Query("comment_id")
	user, _ := c.Get("userinfo")
	temp := user.(*model.Userinfo)
	userinfo := UserLoginInfo{
		Id:            temp.ID,
		Name:          temp.Username,
		FollowCount:   temp.FollowCount,
		FollowerCount: temp.FollowerCount,
		IsFollow:      false,
	}
	comment := model.Comment{Create_date: time.Now().Format("01-02"),
		Content: comment_text, Videoid: uint(video_id), Authorid: temp.ID}
	if action_type == "1" {
		model.Connect2sql().Create(&comment)
		c.JSON(http.StatusOK, struct {
			Response
			Comment CommentResponse
		}{Response: Response{StatusCode: 1, StatusMsg: "发表成功"},
			Comment: CommentResponse{Content: comment_text, User: userinfo, Create_date: comment.Create_date, ID: comment.Id},
		})
	} else if action_type == "2" {
		model.Connect2sql().Table("comments").Where("id = ?", comment_id).Delete(&Comment{})
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "删除成功"})
	}

}

func Commentlist(c *gin.Context) {
	video_id := c.Query("video_id")
	var commentlist []model.Comment
	model.Connect2sql().Where("videoid = ?", video_id).Find(&commentlist)
	lenlist := len(commentlist)
	fmt.Println(lenlist)
	responselist := make([]CommentResponse, lenlist)
	var userinfo model.Userinfo
	loginuser, exist := c.Get("userinfo")

	for index, val := range commentlist {
		responselist[index].ID = val.Id
		responselist[index].Content = val.Content
		responselist[index].Create_date = val.Create_date
		model.Connect2sql().Where("id = ?", val.Authorid).First(&userinfo)
		responselist[index].User = UserLoginInfo{Id: val.Authorid, Name: userinfo.Username, FollowCount: userinfo.FollowerCount,
			FollowerCount: userinfo.FollowerCount, IsFollow: exist && model.Isfollow(loginuser.(*model.Userinfo).ID, val.Authorid)}
		userinfo = model.Userinfo{}
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: responselist,
	})

}
