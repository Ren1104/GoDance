package service

import "douyin/config"

type Following struct {
	// gorm.Model
	Id       int64 `gorm:"column:relation_id; primary_key;"`
	AuthorId int64 `gorm:"column:author_id"`   //视频发布者id
	Follower int64 `gorm:"column:follower_id"` //关注者id
}

func isFollowing(AuthorId int64, Follower int64) bool {
	db := config.DB
	fol := Following{}
	if err := db.Table("following").Where("author_id=? AND follower_id=?", AuthorId, Follower).Find(&fol).Error; err != nil {
		//没有关注
		return false
	}
	//已关注
	return true
}
