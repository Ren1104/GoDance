package service

import (
	"douyin/common"
	"douyin/config"
)

// gorm.Model
type Video struct {
	//gorm.Model
	Id            int64  `gorm:"column:video_id; primary_key;"`
	AuthorId      int64  `gorm:"column:author_id;"`
	PlayUrl       string `gorm:"column:play_url;"`
	CoverUrl      string `gorm:"column:cover_url;"`
	CreateTime    int64  `gorm:"column:create_at;"`
	FavoriteCount int64  `gorm:"column:favorite_count;"`
	CommentCount  int64  `gorm:"column:comment_count;"`
	Title         string `gorm:"column:title;"`
}

const videoNum = 1 //单次最多返回的视频数量

// FeedGet 获得视频列表
func FeedGet(currentTime int64, userId int64) (feedVideoList []common.Video, nextTime int64, err error) {
	var videoList []Video
	videoList = make([]Video, 0)
	videoList, err = GetFeedVideoList(currentTime)
	if err != nil {
		//从头开始
		println("获取视频列表出错")
		println(err.Error())
		//return nil, 0, err
	}

	feedVideoList, nextTime = FeedResponseVideoList(videoList, userId)

	return feedVideoList, nextTime, nil
}

func GetFeedVideoList(currentTime int64) (videoList []Video, err error) {
	db := config.DB
	//按照发布时间获取视频
	err = db.Table("videos").Where("create_at < ?", currentTime).Order("create_at desc").Limit(videoNum).Find(&videoList).Error
	return videoList, err
}

// 把gorm操作的结构体转换成视频流推送的结构体
func FeedResponseVideoList(videoList []Video, userId int64) ([]common.Video, int64) {
	feedVideoList := make([]common.Video, len(videoList))
	var nextTime int64
	for i, v := range videoList {
		author, _ := GetUserInfo(v.AuthorId)
		temp := common.Video{
			Id:            v.Id,
			Author:        FeedResponseUser(author),
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			Title:         v.Title,
		}
		//判断是否关注
		//if ok := isFollowing(v.AuthorId, userId); ok {
		//	temp.Author.IsFollow = true
		//}

		//判断是否点赞
		//if ok := isFavorite(userId, v.Id); ok {
		//	temp.Author.IsFollow = true
		//}
		feedVideoList[i] = temp

		//获取最后的时间
		nextTime = v.CreateTime
	}
	return feedVideoList, nextTime
}
