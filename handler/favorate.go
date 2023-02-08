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
	userinfo, _ := c.Get("userinfo")
	vid, _ := strconv.ParseUint(videoid, 10, 32)
	//连接数据库
	db := model.Connect2sql()
	favInfo := model.Favoriteinfo{Userid: userinfo.(*model.Userinfo).ID, Videoid: uint(vid)}
	res := db.Find(&favInfo)
	if actiontype == "1" {
		if res.RowsAffected < 1 {
			db.Create(&favInfo)
		} else {
			res.Update("valid", true)
		}
	} else if actiontype == "2" {
		if res.RowsAffected > 0 {
			res.Update("valid", false)
		}
	}
	db.First(&model.Videoinfo{}, videoid).Update("Favoritecount", model.Getfovnum(videoid))
	c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: ""})
}

// FavList 喜欢列表
//func FavList(c *gin.Context) {
//	userid := c.Query("user_id")
//
//}
