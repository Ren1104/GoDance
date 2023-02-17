package service

import (
	"GoDance/model"
	"fmt"
)

const videoNum = 2 //单次最多返回的视频数量

// FeedGet 获得视频列表
func FeedGet(currentTime int64, userId int64) (feedVideoList []model.Video, nextTime int64, err error) {
	var videoList []model.VideoData
	videoList = make([]model.VideoData, 0)
	videoList, err = GetFeedVideoList(currentTime)
	if err != nil {
		fmt.Println("获取视频列表出错")
		fmt.Println(err.Error())
		return nil, 0, err
	}

	feedVideoList, nextTime = FeedResponseVideoList(videoList, userId)
	fmt.Println(feedVideoList)
	return feedVideoList, nextTime, nil
}

func GetFeedVideoList(currentTime int64) (videoList []model.VideoData, err error) {
	//按照发布时间获取视频
	err = model.Db.Table("video_data").Where("create_at < ?", currentTime).Order("create_at desc").Limit(videoNum).Find(&videoList).Error
	return videoList, err
}

// 把gorm操作的结构体转换成视频流推送的结构体
func FeedResponseVideoList(videoList []model.VideoData, userId int64) ([]model.Video, int64) {
	feedVideoList := make([]model.Video, len(videoList))
	var nextTime int64
	for i, v := range videoList {
		author, _ := FindUserById(v.Author)
		temp := model.Video{
			Id:            v.Id,
			Author:        author,
			PlayUrl:       v.PlayUrl,
			CoverUrl:      v.CoverUrl,
			FavoriteCount: v.FavoriteCount,
			CommentCount:  v.CommentCount,
			IsFavorite:    false,
			//Title:         v.Title,
			//Author:        FeedResponseUser(author),
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
	fmt.Println(feedVideoList)
	fmt.Println(nextTime)
	return feedVideoList, nextTime
}
