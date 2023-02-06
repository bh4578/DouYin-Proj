package model

import "gorm.io/gorm"

type Videoinfo struct {
	gorm.Model
	Authorid      int
	Playurl       string
	Coverurl      string
	Favoritecount int
	Commentcount  int
	Title         string
}

type Favoriteinfo struct {
	gorm.Model
	Userid  uint64
	Videoid uint64
}

type Commentinfo struct {
	gorm.Model
	Userid  uint64
	Videoid uint64
	Content string
}
