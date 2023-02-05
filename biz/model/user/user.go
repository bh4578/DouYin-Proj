package user

import "gorm.io/gorm"

type Userinfo struct {
	gorm.Model
	//用户名与密码
	Username string `gorm:"unique"`
	Password string
	//关注数与粉丝数
	FollowCount   int
	FollowerCount int
	//已关注用ID
	FollowerID string
	//权限,默认为1
	Authority int `gorm:"default:1"`
}

func (u Userinfo) TableName() string {
	return "userinfo"
}
