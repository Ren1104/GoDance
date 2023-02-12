package service

import (
	"douyin/common"
	"douyin/config"
)

type User struct {
	// gorm.Model
	Id            int64  `gorm:"column:user_id; primary_key;"`
	Name          string `json:"name"`
	Password      string `json:"password"`
	FollowCount   int64  `json:"follow_count"`
	FollowerCount int64  `json:"follower_count"`
	TotalFav      int64  `json:"total_favorited"`
	FavCount      int64  `json:"favorite_count"`
}

func GetUserInfo(userId int64) (User, error) {
	db := config.DB
	user := User{}
	if err := db.Table("users").Where("id=  ?", userId).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func FeedResponseUser(user User) common.User {
	return common.User{
		Id:            user.Id,
		Name:          user.Name,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      false,
		TotalFav:      user.TotalFav,
		FavCount:      user.FavCount,
	}
}
