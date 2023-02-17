package controller

import (
	"GoDance/model"
	"GoDance/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoId := c.Query("video_id")
	actionType := c.Query("action_type")
	_, err := service.FindUserByToken(token)
	if err == nil {
		if actionType == "2" { //delete the favorite
			favorite := model.FavoriteData{Token: token, VideoId: videoId}
			model.Db.Table("favorite_data").Where("token=?", token).Delete(&favorite)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else if actionType == "1" {
			favorite := model.FavoriteData{Token: token, VideoId: videoId}
			model.Db.Create(&favorite)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "The entry doesn't exist"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	var video_list []model.VideoData
	token := c.Query("token")
	model.Db.Table("favorite_data").Find(&video_list, "token=?", token)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		//VideoList: video_list,
	})
}
