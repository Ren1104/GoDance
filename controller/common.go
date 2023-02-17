package controller
import (
	"GoDance/model"
)
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty" grom:"autoIncrement"`
	//Author        model.User   `json:"author" gorm:"foreignKey:Id;references:Id;"`
	Author int64 `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty" gorm:"default:https://www.w3schools.com/html/movie.mp4"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"default:https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg"`
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"default:0"`
	CommentCount  int64  `json:"comment_count,omitempty" gorm:"default:0"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"type:tinyint(1);default:0"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       model.User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

// type User struct {
// 	Id            int64  `json:"id,omitempty"`
// 	Name          string `json:"name,omitempty"`
// 	FollowCount   int64  `json:"follow_count,omitempty"`
// 	FollowerCount int64  `json:"follower_count,omitempty"`
// 	IsFollow      bool   `json:"is_follow,omitempty"`
// }

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

