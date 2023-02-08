package handler

import (
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"path"
	"strconv"
	"time"
)

func Publish(c *gin.Context) {

	file, _ := c.FormFile("data")
	author, _ := c.Get("userinfo")
	title := c.PostForm("title")
	ip := "http://192.168.23.4:8080/static/"
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) +
		author.(*model.Userinfo).Username + "_" + file.Filename
	filepath := path.Join("./public/", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	videoinfo := model.Videoinfo{Playurl: ip + filename, Authorid: author.(*model.Userinfo).ID, Title: title,
		Coverurl: "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"}
	model.Connect2sql().Create(&videoinfo)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

func Publishlist(c *gin.Context) {
	var videolist []model.Videoinfo
	userid := c.Query("user_id")
	db := model.Connect2sql()
	db.Where("authorid = ?", userid).Order("created_at desc").Find(&videolist)
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
			resvideos[index].IsFavorite = model.Isfavorite(userinfo[index].ID, val.ID)
			resvideos[index].Title = val.Title
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: resvideos,
		})

	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "加载喜欢列表错误"})
	}
}
