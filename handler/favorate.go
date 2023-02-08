package handler

import (
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type FavListResponse struct {
	Response          //返回状态信息
	VideoList []Video `json:"video_list"` // 用户点赞视频列表
}

// FavAction 赞操作
func FavAction(c *gin.Context) {
	videoid := c.Query("video_id")
	actiontype := c.Query("action_type") //1-点赞，2-取消点赞

	//获取userid
	_, userinfo := model.Checktoken(c)
	userid := userinfo.ID
	//连接数据库
	db := model.Connect2sql()
	videoidnum, _ := strconv.Atoi(videoid)
	favInfo := model.Favoriteinfo{Userid: uint64(userid), Videoid: uint64(videoidnum)}
	if actiontype == "1" {
		db.Create(&favInfo)
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "点赞成功"})
	} else {
		db.Delete(&favInfo)
		c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "取消点赞"})
	}
}

// FavList 喜欢列表
//func FavList(c *gin.Context) {
//	userid := c.Query("user_id")
//
//}
