package model

type Following struct {
	Id       int64 `gorm:"column:relation_id; primary_key;"`
	AuthorId int64 `gorm:"column:author_id"`   //视频发布者id
	Follower int64 `gorm:"column:follower_id"` //关注者id
}
