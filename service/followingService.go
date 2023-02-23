package service

import (
	"GoDance/model"
)

func isFollowing(AuthorId int64, Follower int64) bool {
	fol := model.Following{}
	if err := model.Db.Table("following").Where("author_id=? AND follower_id=?", AuthorId, Follower).Find(&fol).Error; err != nil {
		//没有关注
		return false
	}
	//已关注
	return true
}
