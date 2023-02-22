package controller

import (
	"GoDance/model"
	"GoDance/service"
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)
//static resource dictionary of the server
var ResourceDir = "http://192.168.115.80:8080/douyin/public/"

// convert video into cover using ffmpeg
func Video2Cover(VideoName string) (CoverPath string, err error) {
	VideoPath := filepath.Join(".\\public\\videos", VideoName)
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(VideoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,1)")}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	cover, err := imaging.Decode(buf)
	if err != nil {
		return "", err
	}
	CoverPath = ".\\public\\covers\\" + VideoName + ".png"
	//fmt.Println("CoverPath:  "+ CoverPath)
	err = imaging.Save(cover, CoverPath)
	if err != nil {
		fmt.Printf("Saving cover error\n %s\n", err)
		return "", err
	}
	return CoverPath, nil
}

// Publish !! Title didn't use !!
// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	if model.Err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  model.Err.Error(),
		})
		fmt.Printf("database error\n")
		return
	}
	token := c.PostForm("token")

	login_user, err := service.FindUserByToken(token)
	//fmt.Println(login_user)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//fmt.Println(login_user.Id)
	filename := filepath.Base(data.Filename)
	finalName := fmt.Sprintf("%d_%s", login_user.Id, filename)
	saveFile := filepath.Join("./public/videos", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	//title := c.PostForm("title")
	VideoUrl := filepath.Join(ResourceDir, "videos", finalName)
	CoverUrl, err := Video2Cover(finalName)
	CoverUrl = filepath.Join(ResourceDir, "covers", finalName+".png")
	//fmt.Printf("video: %s    cover: %s\n", VideoUrl, CoverUrl)
	new_video := model.VideoData{
		Author:        login_user.Id,
		PlayUrl:       VideoUrl,
		CoverUrl:      CoverUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		CreateTime:    time.Now().Unix(),
	}
	model.Db.Create(&new_video)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	VideoList := []model.VideoData{}

	author_id := c.Query("user_id")
	model.Db.Table("video_data").Find(&VideoList, "author=?", author_id)
	//search author row
	authorid, _ := strconv.ParseInt(author_id, 10, 64)
	author, err := service.FindUserById(authorid)
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Author doesn't exist"})
		return
	}

	var PublishVideoList []model.PublishListVideoStruct //[]PublishListVideoStruct{} //video list used for publish list interface
	for i := 0; i < len(VideoList); i++ {
		var tmp = model.PublishListVideoStruct{
			Id:            VideoList[i].Id,
			Author:        author,
			PlayUrl:       VideoList[i].PlayUrl,
			CoverUrl:      VideoList[i].CoverUrl,
			FavoriteCount: VideoList[i].FavoriteCount,
			CommentCount:  VideoList[i].CommentCount,
			IsFavorite:    VideoList[i].IsFavorite,
		}
		PublishVideoList = append(PublishVideoList, tmp)
	}

	c.JSON(http.StatusOK, PublishVideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		PublishVideoList: PublishVideoList,
	})
}
