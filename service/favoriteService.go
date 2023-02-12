package service

import "douyin/config"

type Favorite struct {
	Id      int64 `gorm:"column:favorite_id; primary_key;"`
	UserId  int64 `gorm:"column:user_id"`  //用户
	VideoId int64 `gorm:"column:video_id"` //喜欢的视频id
}

func isFavorite(UserId int64, VideoId int64) bool {
	db := config.DB
	fol := Favorite{}
	if err := db.Table("favorite").Where("user_id=? AND guest_id=?", UserId, VideoId).Find(&fol).Error; err != nil {
		//没点喜欢
		return false
	}
	//点了喜欢
	return true
}
