package model

import "time"

// 给朋友发私信
type Message struct {
	Id          uint `gorm:"primaryKey"`
	Authorid    uint
	Targetid    uint
	Content     string
	Create_date time.Time
}

// 用户评论
type Comment struct {
	Id          uint `gorm:"primaryKey"`
	Authorid    uint
	Videoid     uint
	Content     string
	Create_date string
}
