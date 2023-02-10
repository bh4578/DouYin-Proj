package handler

// 该包主要用户放置一些公用的结构体
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type Video struct {
	Id            uint          `json:"id,omitempty"`
	Author        UserLoginInfo `json:"author"`
	PlayUrl       string        `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string        `json:"cover_url,omitempty"`
	FavoriteCount int           `json:"favorite_count,omitempty"`
	CommentCount  int           `json:"comment_count,omitempty"`
	IsFavorite    bool          `json:"is_favorite,omitempty"`
	Title         string        `json:"title,omitempty"`
}

type Comment struct {
	Id         uint          `json:"id,omitempty"`
	User       UserLoginInfo `json:"user_router"`
	Content    string        `json:"content,omitempty"`
	CreateDate string        `json:"create_date,omitempty"`
}

type UserLoginInfo struct {
	Id            uint   `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int    `json:"follow_count,omitempty"`
	FollowerCount int    `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type MessageSendEvent struct {
	UserId     uint   `json:"user_id,omitempty"`
	ToUserId   int    `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId uint   `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}
