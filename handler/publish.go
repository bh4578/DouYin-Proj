package handler

import (
	"Douyin/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
	"path"
	"strconv"
	"strings"
	"time"
)

func getvideocover(filepath string) {
	coverPath := strings.Split(filepath, ".")[0] + ".jpg"
	cmd := exec.Command("ffmpeg", "-i", filepath, "-vframes", "1", "-ss", "00:00:01", "-f", "image2", coverPath)
	cmd.Run()
	return
}

func Publish(c *gin.Context) {

	file, _ := c.FormFile("data")
	author, _ := c.Get("userinfo")
	title := c.PostForm("title")
	ip := "http://192.168.23.60:8080/static/"
	temp := strings.Split(file.Filename, ".")
	suffix := "." + temp[len(temp)-1]
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) +
		author.(*model.Userinfo).Username + suffix
	filepath := path.Join("./public/", filename)
	if err := c.SaveUploadedFile(file, filepath); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	covername := ip + strings.Split(filename, ".")[0] + ".jpg"
	getvideocover(filepath)
	videoinfo := model.Videoinfo{Playurl: ip + filename, Authorid: author.(*model.Userinfo).ID, Title: title,
		Coverurl: covername}
	model.Connect2sql().Create(&videoinfo)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  "uploaded successfully",
	})
}

func Publishlist(c *gin.Context) {
	var videolist []model.Videoinfo
	userid := c.Query("user_id")
	loginuser, exist := c.Get("userinfo")
	db := model.Connect2sql()
	db.Where("authorid = ?", userid).Order("created_at desc").Find(&videolist)
	lenoflist := len(videolist)
	var userinfo model.Userinfo

	if lenoflist > 0 {
		resvideos := make([]Video, lenoflist)
		for index, val := range videolist {
			resvideos[index].Id = val.ID
			db.Where("id = ?", val.Authorid).First(&userinfo)
			resvideos[index].Author = UserLoginInfo{Id: userinfo.ID, Name: userinfo.Username,
				FollowCount: userinfo.FollowCount, FollowerCount: userinfo.FollowerCount,
				IsFollow: exist && model.Isfollow(loginuser.(*model.Userinfo).ID, val.Authorid)}
			resvideos[index].CommentCount = val.Commentcount
			resvideos[index].FavoriteCount = val.Favoritecount
			resvideos[index].CoverUrl = val.Coverurl
			resvideos[index].PlayUrl = val.Playurl
			resvideos[index].IsFavorite = model.Isfavorite(userinfo.ID, val.ID)
			resvideos[index].Title = val.Title
			userinfo = model.Userinfo{}
		}
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0},
			VideoList: resvideos,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "加载发布列表错误"})
	}
}
