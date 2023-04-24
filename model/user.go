package model

import "gorm.io/gorm"

// 用户信息表
type Userinfo struct {
	gorm.Model
	//用户名与密码
	Username string `gorm:"unique"`
	Password string
	//关注数与粉丝数
	FollowCount   int
	FollowerCount int
	//权限,默认为1，不需要！！！！本意禁言
	Authority int `gorm:"default:1"`
}

// 用户关注关系表
type Userrelation struct {
	gorm.Model
	Userid   uint
	Targetid uint
	Valid    bool `gorm:"default:true"`
}
