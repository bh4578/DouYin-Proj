package model

import "time"

type Message struct {
	Id          uint `gorm:"primaryKey"`
	Authorid    uint
	Targetid    uint
	Content     string
	Create_date time.Time
}
type Comment struct {
	Id          uint `gorm:"primaryKey"`
	Authorid    uint
	Videoid     uint
	Content     string
	Create_date string
}
