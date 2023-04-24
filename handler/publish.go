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

// 获取视频第一帧
func getvideocover(filepath string) {
	coverPath := strings.Split(filepath, ".")[0] + ".jpg"
	cmd := exec.Command("ffmpeg", "-i", filepath, "-vframes", "1", "-ss", "00:00:01", "-f", "image2", coverPath)
	cmd.Run()
	return
}

func Publish(c *gin.Context) {
	// 用于获取HTTP请求中上传的文件
	// "name"参数时上传的文件字段名，返回值是一个*multipart.FileHeader 类型指针和error
	// 返回的文件头 *multipart.FileHeader 对象，获取到上传文件的各种信息。
	file, _ := c.FormFile("data")
	/*
		Get获取的值返回的是一个接口类型的值，因为 gin.Context 用于处理 HTTP 请求，而请求中传递的参数类型各不相同
		因此 gin.Context 的 Get() 方法返回的值也需要支持各种类型，故使用了接口类型。
		通过断言将 author 这个接口类型的值转换成了 *model.Userinfo 类型，以便于访问其中的字段。
		因为接口类型可以存储任何类型的值，但是我们需要使用其中的具体字段和方法，需要将其转换为其原始类型，而这个过程被称为类型断言。
	*/
	author, _ := c.Get("userinfo")
	// 用于获取POST请求中form表单数据的方法
	// 如果POST请求表中有一个title的值，则返回该值，不存在返回空字符串
	title := c.PostForm("title")
	ip := "http://192.168.23.60:8080/static/"
	temp := strings.Split(file.Filename, ".")
	suffix := "." + temp[len(temp)-1]
	filename := strconv.FormatInt(time.Now().UnixNano(), 10) +
		author.(*model.Userinfo).Username + suffix
	filepath := path.Join("./public/", filename)
	// 将用户上传视频文件保存到服务器本地
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

// 获取指定用户发布的视频列表
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
