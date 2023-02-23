package model

// 存储在数据库中的User结构
type UserData struct {
	Id            int64  `gorm:"column:user_id; primary_key;""`
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	TotalFav      int64  `json:"total_favorited"`
	FavCount      int64  `json:"favorite_count"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Token         string `json:"token,omitempty"`
}

// 用于返回的User结构
type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Password      string `json:"password,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	TotalFav      int64  `json:"total_favorited,omitempty"`
	FavCount      int64  `json:"favorite_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
	Token         string `json:"token,omitempty"`
}
