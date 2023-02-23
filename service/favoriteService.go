package service

import (
	"GoDance/model"
)

func isFavorite(UserId int64, VideoId int64) bool {
	fol := model.FavoriteData{}
	if err := model.Db.Table("favorite").Where("user_id=? AND guest_id=?", UserId, VideoId).Find(&fol).Error; err != nil {
		//没点喜欢
		return false
	}
	//点了喜欢
	return true
}
