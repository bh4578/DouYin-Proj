package handler

import (
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video_router list for every request
func Feed(c *gin.Context) {
	lasttime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	var videolist []model.Videoinfo
	db := model.Connect2sql()
	db.Where("created_at > ?", time.Unix(lasttime, 0).
		Format("2006-01-02 15:04:05.000")).Limit(30).Order("created_at desc").Find(&videolist)
	lenoflist := len(videolist)

	var userinfo model.Userinfo
	flag, loginuser := model.Checktoken(c)
	if lenoflist > 0 {
		resvideos := make([]Video, lenoflist)
		for index, val := range videolist {
			resvideos[index].Id = val.ID
			db.Where("id = ?", val.Authorid).First(&userinfo)
			resvideos[index].Author = UserLoginInfo{Id: userinfo.ID, Name: userinfo.Username, FollowCount: userinfo.FollowCount, FollowerCount: userinfo.FollowerCount, IsFollow: flag && model.Isfollow(loginuser.ID, val.Authorid)}
			resvideos[index].CommentCount = val.Commentcount
			resvideos[index].FavoriteCount = val.Favoritecount
			resvideos[index].CoverUrl = val.Coverurl
			resvideos[index].PlayUrl = val.Playurl
			resvideos[index].IsFavorite = flag && model.Isfavorite(loginuser.ID, val.ID)
			resvideos[index].Title = val.Title
			userinfo = model.Userinfo{}
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
