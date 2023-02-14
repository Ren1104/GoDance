package model

import (
	"GoDance/controller"
	"time"
)

type CommentData struct {
	Id          int64
	UserId      int64
	VideoId     int64
	CommentText string
	CreateDate  time.Time
	Cancel      int8
}

func InsertComment(comment *CommentData) error {
	if err := controller.DB.Create(comment).Error; err != nil {
		return err //添加评论失败
	}
	return nil //添加评论成功
}

func DeleteComment(id int64) error {
	//cancel字段默认0表示存在,1表示逻辑删除
	if err := controller.DB.Model(&CommentData{}).Where("id=?", id).Update("cancel", 1).Error; err != nil {
		return err //删除失败
	} else {
		return nil //删除成功
	}
}

func FindCommentById(id int64) (CommentData, error) {
	var comment CommentData
	if err := controller.DB.Where(&CommentData{Id: id}).First(&comment).Error; err != nil {
		return CommentData{}, err
	}
	return comment, nil

}

func FindCommentsByVideoId(viderId int64) []CommentData {
	var comments []CommentData
	//当使用结构作为条件查询时，GORM 只会查询非零值字段，更新也是如此。
	//config.DB.Where(&Comment{Cancel: 0, VideoId: viderId}).Order("create_date desc").Find(&comments)
	//解决办法：使用 map 来构建更新条件，而不是结构体 &Comment{}
	controller.DB.Where(map[string]interface{}{"video_id": viderId, "cancel": 0}).Order("create_date desc").Find(&comments)
	return comments
}
