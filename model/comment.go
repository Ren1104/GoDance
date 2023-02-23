package model

import (
	"time"

	"gorm.io/gorm"
)

// 用于返回的comment结构
type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

// 存储在数据库中的comment结构
type CommentData struct {
	Id          int64     `json:"id,omitempty"`
	UserId      int64     `json:"userId,omitempty"`
	VideoId     int64     `json:"videoId,omitempty"`
	CommentText string    `json:"commentText,omitempty"`
	CreateDate  time.Time `json:"createDate,omitempty"`
	Cancel      int8      `json:"cancel,omitempty"`
}

func InsertComment(comment *CommentData) error {
	err1 := Db.Transaction(func(db *gorm.DB) error {
		if err := db.Create(comment).Error; err != nil {
			return err
		}
		if err := db.Model(&VideoData{}).Where("Id = ?", comment.VideoId).Update("CommentCount", gorm.Expr("comment_count + ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err1

}

func DeleteComment(id int64, videoid int64) error {
	err1 := Db.Transaction(func(db *gorm.DB) error {

		if err := db.Model(&CommentData{}).Where("Id = ?", id).Update("cancel", 1).Error; err != nil {
			return err
		}

		if err := db.Model(&VideoData{}).Where("Id = ?", videoid).Update("CommentCount", gorm.Expr("comment_count - ?", 1)).Error; err != nil {
			return err
		}
		return nil
	})
	return err1
}

func FindCommentById(id int64) (CommentData, error) {
	var comment CommentData
	if err := Db.Where(&CommentData{Id: id}).First(&comment).Error; err != nil {
		return CommentData{}, err
	}
	return comment, nil

}

func FindCommentsByVideoId(viderId int64) []CommentData {
	var comments []CommentData
	//当使用结构作为条件查询时，GORM 只会查询非零值字段，更新也是如此。
	//config.DB.Where(&Comment{Cancel: 0, VideoId: viderId}).Order("create_date desc").Find(&comments)
	//解决办法：使用 map 来构建更新条件，而不是结构体 &Comment{}
	Db.Where(map[string]interface{}{"video_id": viderId, "cancel": 0}).Order("create_date desc").Find(&comments)
	return comments
}
