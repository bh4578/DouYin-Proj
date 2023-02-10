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
	videoid_int, _ := strconv.ParseInt(videoid, 10, 64)
	//连接数据库
	db := model.Connect2sql()
	favInfo := model.Favoriteinfo{}
	res := db.Where("userid = ? AND videoid = ?", userinfo.(*model.Userinfo).ID, videoid_int).Find(&favInfo)
	if actiontype == "1" {
		if res.RowsAffected < 1 {
			favInfo.Videoid = uint(videoid_int)
			favInfo.Userid = userinfo.(*model.Userinfo).ID
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
func FavList(c *gin.Context) {
	userid := c.Query("user_id")
	var videolist []model.Videoinfo
	db := model.Connect2sql()
	db.Table("videoinfos").Joins("inner join favoriteinfos on videoinfos.id = favoriteinfos.videoid").Where("favoriteinfos.userid = ? AND favoriteinfos.valid = ?", userid, true).Scan(&videolist)
	loginid, _ := c.Get("userinfo")
	lenoflist := len(videolist)
	userinfo := make([]model.Userinfo, lenoflist)
	if lenoflist > 0 {
		resvideos := make([]Video, lenoflist)
		for index, val := range videolist {
			resvideos[index].Id = val.ID
			db.Where("id = ?", val.Authorid).First(&userinfo[index])
			resvideos[index].Author = UserLoginInfo{Id: userinfo[index].ID, Name: userinfo[index].Username, FollowCount: 0, FollowerCount: 0}
			resvideos[index].CommentCount = val.Commentcount
			resvideos[index].FavoriteCount = val.Favoritecount
			resvideos[index].CoverUrl = val.Coverurl
			resvideos[index].PlayUrl = val.Playurl
			resvideos[index].IsFavorite = model.Isfavorite(loginid.(*model.Userinfo).ID, val.ID)
			resvideos[index].Title = val.Title
		}

		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: resvideos,
			NextTime:  videolist[0].CreatedAt.Unix(),
		})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "请重新刷新列表"})
	}

}
