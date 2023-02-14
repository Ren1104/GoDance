package model

import "GoDance/controller"

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

func GetUserInfoVO(userId int64) (controller.User, error) {
	db := controller.DB
	user := User{}
	var userVO controller.User
	if err := db.Table("users").Where("id=  ?", userId).Find(&user).Error; err != nil {
		return userVO, err
	}
	userVO.FollowCount = user.FollowCount
	userVO.Id = user.Id
	userVO.Name = user.Name
	userVO.FollowerCount = user.FollowerCount
	userVO.IsFollow = userVO.IsFollow
	return userVO, nil
}
