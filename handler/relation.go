package handler

import (
	"Douyin/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Follow(c *gin.Context) {
	to_user_id, _ := strconv.ParseInt(c.Query("to_user_id"), 10, 64)
	actiontype := c.Query("action_type")
	loginuser, _ := c.Get("userinfo")
	userid := loginuser.(*model.Userinfo).ID

	if uint(to_user_id) == userid {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "不能自己关注自己！"})
		return
	}
	db := model.Connect2sql()
	var count int64
	res := db.Table("userrelations").Where("userid = ? AND targetid = ?", userid, to_user_id).Count(&count)
	//var userrelation model.Userrelation
	if actiontype == "1" {
		if count < 1 {
			db.Create(&model.Userrelation{Userid: userid, Targetid: uint(to_user_id)})
		} else {
			////因为app有bug，所以后台多查询一次，判断是否重复关注
			//res.First(&userrelation)
			//if userrelation.Valid {
			//	c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "请勿重复关注!"})
			//	return
			//}
			res.Update("valid", true)
		}
	} else if actiontype == "2" {
		if count > 0 {
			res.Update("valid", false)
		}
	}
	db.First(&model.Userinfo{}, userid).Update("follow_count", model.Getfollownum(userid))
	db.First(&model.Userinfo{}, to_user_id).Update("follower_count", model.Getfollowernum(uint(to_user_id)))
	c.JSON(http.StatusOK, Response{StatusCode: 0})
}

//关注列表

type Followlist struct {
	Response
	User_list []UserLoginInfo `json:"user_list"`
}

func Getfollowlist(c *gin.Context) {

	user_id := c.Query("user_id")
	var userlist []model.Userinfo
	db := model.Connect2sql()
	loginuser, exist := c.Get("userinfo")
	db.Table("userinfos").Select("userinfos.id,username,follow_count,follower_count").Joins("inner join userrelations on userinfos.id = userrelations.targetid").Where("userrelations.userid = ? AND valid = ?", user_id, true).Scan(&userlist)
	reslist := make([]UserLoginInfo, len(userlist))
	for index, val := range userlist {
		reslist[index].Id = val.ID
		reslist[index].FollowerCount = val.FollowerCount
		reslist[index].FollowCount = val.FollowCount
		reslist[index].Name = val.Username
		reslist[index].IsFollow = exist && model.Isfollow(loginuser.(*model.Userinfo).ID, val.ID)
	}

	c.JSON(http.StatusOK, Followlist{Response: Response{StatusCode: 0}, User_list: reslist})
}

//粉丝列表

func Getfollowerlist(c *gin.Context) {

	user_id := c.Query("user_id")
	var userlist []model.Userinfo
	db := model.Connect2sql()
	loginuser, exist := c.Get("userinfo")
	db.Table("userinfos").Select("userinfos.id,username,follow_count,follower_count").Joins("inner join userrelations on userinfos.id = userrelations.userid").Where("userrelations.targetid = ? AND valid = ?", user_id, true).Scan(&userlist)

	reslist := make([]UserLoginInfo, len(userlist))
	for index, val := range userlist {
		reslist[index].Id = val.ID
		reslist[index].FollowerCount = val.FollowerCount
		reslist[index].FollowCount = val.FollowCount
		reslist[index].Name = val.Username
		reslist[index].IsFollow = exist && model.Isfollow(loginuser.(*model.Userinfo).ID, val.ID)
	}
	fmt.Println(loginuser.(*model.Userinfo).ID)
	fmt.Printf("%+v", reslist)
	fmt.Println()

	c.JSON(http.StatusOK, Followlist{Response: Response{StatusCode: 0}, User_list: reslist})
}

func Getfriendlist(c *gin.Context) {

	user_id := c.Query("user_id")
	var userlist []model.Userinfo
	db := model.Connect2sql()
	loginuser, exist := c.Get("userinfo")
	db.Table("userinfos").
		Select("id,username,follow_count,follower_count").
		Joins("JOIN (SELECT a.userid FROM userrelations a JOIN userrelations b ON a.userid = b.targetid AND a.targetid = b.userid WHERE a.valid = true AND b.valid = true AND a.targetid = ?) userrelation ON userinfos.id = userrelation.userid", user_id).
		Scan(&userlist)
	reslist := make([]UserLoginInfo, len(userlist))
	for index, val := range userlist {
		reslist[index].Id = val.ID
		reslist[index].FollowerCount = val.FollowerCount
		reslist[index].FollowCount = val.FollowCount
		reslist[index].Name = val.Username
		reslist[index].IsFollow = exist && model.Isfollow(loginuser.(*model.Userinfo).ID, val.ID)
	}

	c.JSON(http.StatusOK, Followlist{Response: Response{StatusCode: 0}, User_list: reslist})

}
