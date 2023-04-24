package handler

import (
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv" //字符串和基本数据类型之间相互转换
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video_router list for every request
func Feed(c *gin.Context) {
	// base:进制        bitSize：int64
	lasttime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	var videolist []model.Videoinfo
	db := model.Connect2sql()
	// time.Unix:转换为事件对象      Format:返回时间值的文本表示形式
	// Limit(30) 限制查询结果数量30   Order:指定从数据库中检索记录时的顺序，按照created_at字段倒序排列
	db.Where("created_at > ?", time.Unix(lasttime, 0).
		Format("2006-01-02 15:04:05.000")).Limit(30).Order("created_at desc").Find(&videolist)
	lenoflist := len(videolist)

	var userinfo model.Userinfo
	flag, loginuser := model.Checktoken(c)
	if lenoflist > 0 {
		// 一个[]Video类型，长度lenflist的切片
		resvideos := make([]Video, lenoflist)
		// 循环将数据库中查询的videolist信息传给resvideos
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
		// FeedResponse结构体中包含[]Video类型，即resvideos。
		// Json格式响应返回
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: resvideos,
			NextTime:  videolist[0].CreatedAt.Unix(),
		})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "请重新刷新列表"})
	}

}
