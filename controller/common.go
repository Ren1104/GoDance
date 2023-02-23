package controller

import "GoDance/model"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// Comment About comment

type CommentListResponse struct {
	Response
	CommentList []model.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment model.Comment `json:"comment,omitempty"`
}

// Video About Video

type FeedResponse struct {
	Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []model.Video `json:"video_list"`
}
type PublishVideoListResponse struct {
	Response
	PublishVideoList []model.PublishListVideoStruct `json:"video_list"`
}

// User About User

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type LoginResponse struct {
	Response
	Id    int64  `json:"user_id,omitempty"`
	Token string `json:"token,omitempty"`
}
