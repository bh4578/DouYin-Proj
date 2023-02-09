package model

type Message struct {
	Id          uint `gorm:"primaryKey"`
	Content     string
	Create_date string
}
type Comment struct {
	Id          uint `gorm:"primaryKey"`
	Authorid    uint
	Videoid     uint
	Content     string
	Create_date string
}
