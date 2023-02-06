package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

// Feed same demo video_router list for every request
func Feed(c *gin.Context) {
	//lasttime, _ := strconv.ParseInt(c.Query("latest_time"), 10, 64)
	var DemoVideos = []Video{
		{
			Id: 1,
			Author: UserLoginInfo{
				Id:            1,
				Name:          "TestUser",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false},
			PlayUrl:       "http://192.168.23.183:8080/static/testmv.mp4",
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		},
		{
			Id: 2,
			Author: UserLoginInfo{
				Id:            1,
				Name:          "TestUser",
				FollowCount:   0,
				FollowerCount: 0,
				IsFollow:      false},
			PlayUrl:       "http://192.168.23.183:8080/static/bear.mp4",
			CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
			FavoriteCount: 0,
			CommentCount:  0,
			IsFavorite:    false,
		},
	}

	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
