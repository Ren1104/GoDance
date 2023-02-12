package controller

import (
	"douyin/common"
	"douyin/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

//目前controller部分只有fead.go改动过

type FeedResponse struct {
	common.Response
	VideoList []common.Video `json:"video_list,omitempty"`
	NextTime  int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	//token := c.Query("token")
	//userIds, _ := c.Get("UserId")
	//userId := userIds.(int64)
	currentTime, err := strconv.ParseInt(c.Query("latest_time"), 10, 64)

	if err != nil || currentTime == 0 {
		currentTime = time.Now().Unix()
	}

	//FeedGet的参数二是userId,这里测试所以设置为0
	feedList, nextTime, _ := service.FeedGet(currentTime, 0)

	c.JSON(http.StatusOK, FeedResponse{
		Response:  common.Response{StatusCode: 0}, //成功
		VideoList: feedList,
		NextTime:  nextTime,
	})
}

//demo
/*
func Feed(c *gin.Context) {
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	})
}
*/
