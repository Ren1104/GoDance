package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Favorite struct {
	token    string `gorm:"column:token"`
	video_id string `gorm:"column:video_id"`
	//user_id  int64  `gorm:"column:user_id"`	//because FavoriteAction has no req para named user_id
}

func (Favorite) TableName() string {
	return "favorite"
}

var DB *gorm.DB
var err error

func initMysql() {
	//config of Mysql, wait to be adjusted when DB is settled
	dsn := "root:123456@tcp(127.0.0.1:3306)/sky?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
}

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	action_type := c.Query("action_type")

	initMysql()

	if _, exist := usersLoginInfo[token]; exist {
		if action_type == "2" { //delete the favorite
			favorite := Favorite{token: token, video_id: video_id}
			DB.Delete(&favorite)
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		} else if action_type == "1" {
			favorite := Favorite{token: token, video_id: video_id}
			DB.Create(&favorite)
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
	var video_list []Video
	var condition = "token" + c.Query("token")
	DB.Where(condition).Find(&video_list)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: video_list,
	})
}
