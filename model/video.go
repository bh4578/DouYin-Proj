package model

import "gorm.io/gorm"

type Videoinfo struct {
	gorm.Model
	Authorid      uint
	Playurl       string
	Coverurl      string
	Favoritecount int
	Commentcount  int
	Title         string
}

type Favoriteinfo struct {
	gorm.Model
	Userid  uint
	Videoid uint
	Valid   bool `gorm:"default:true"`
}
