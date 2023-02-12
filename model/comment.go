package model

import (
	"GoDance/controller"
	"time"
)

type Comment struct {
	Id          int64
	UserId      int64
	VideoId     int64
	CommentText string
	CreateDate  time.Time
	Cancel      int8
}

func InsertComment(comment *Comment) {
	controller.DB.Create(comment)
}

func DeleteComment(id int64) {
	//cancel字段默认0表示存在,1表示逻辑删除
	controller.DB.Model(&Comment{}).Where("id=?", id).Update("cancel", 1)
}

func FindCommentById(id int64) Comment {
	var comment Comment
	controller.DB.Where(&Comment{Id: id}).First(&comment)
	return comment
}

func FindCommentsByVideoId(viderId int64) []Comment {
	var comments []Comment
	//当使用结构作为条件查询时，GORM 只会查询非零值字段，更新也是如此。
	//config.DB.Where(&Comment{Cancel: 0, VideoId: viderId}).Order("create_date desc").Find(&comments)
	//解决办法：使用 map 来构建更新条件，而不是结构体 &Comment{}
	controller.DB.Where(map[string]interface{}{"video_id": viderId, "cancel": 0}).Order("create_date desc").Find(&comments)
	return comments
}
