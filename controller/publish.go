package controller
import (
	"fmt"
	"GoDance/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/disintegration/imaging"
	"bytes"
	"os"
)
var ResourceDir = "http://127.0.0.1:8080/static/"
type VideoListResponse struct {
	Response
	VideoList []PublishListVideoStruct `json:"video_list"`
}

//as just store author's id in database, need another struct to meet the interface of PublishList
type PublishListVideoStruct struct {
	Id            int64  `json:"id,omitempty"`
	Author        model.User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

//convert video into cover using ffmpeg
func Video2Cover(VideoName string) (CoverPath string, err error) {
	VideoPath := filepath.Join("./public/videos", VideoName)
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(VideoPath).Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,1)")}).
		Output("pipe:",ffmpeg.KwArgs{"vframes":1, "format":"image2", "vcodec":"mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		fmt.Printf("Generating cover error\n")
		return "", err
	}
	cover, err := imaging.Decode(buf)
	if err != nil {
		fmt.Printf("Generating cover error\n")
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

// !! Title didn't use !!
// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})  
		fmt.Printf("database error\n")
		return
	}
	token := c.PostForm("token")

	var login_user model.User
	var res = db.Find(&login_user, "token=?", token)
	if res.Error != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User info invalid"})
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
	fmt.Println(login_user.Id)
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
	new_video := Video{
		Author: login_user.Id,
		PlayUrl: VideoUrl,
		CoverUrl: CoverUrl,
	}
	db.Create(&new_video)
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	VideoList := []Video{}
	
	author_id := c.Query("user_id")
	db.Find(&VideoList, "author=?", author_id)
	
	//search author row
	var author model.User
	var res = db.Find(&author, "id=?", author_id)
	if res.Error != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Author doesn't exist"})
	 	return
	}

	PublishVideoList := []PublishListVideoStruct{} //[]PublishListVideoStruct{} //video list used for publish list interface
	for i := 0; i < len(VideoList); i++ {
		var tmp  = PublishListVideoStruct{
			Id: VideoList[i].Id,
			Author: author,
			PlayUrl: VideoList[i].PlayUrl,
			CoverUrl: VideoList[i].CoverUrl,
			FavoriteCount: VideoList[i].FavoriteCount,
			CommentCount: VideoList[i].CommentCount, 
			IsFavorite: VideoList[i].IsFavorite,
		}
		// fmt.Println(tmp)
		// fmt.Println("!!!")
		PublishVideoList = append(PublishVideoList, tmp)	
	}
	// fmt.Print("Orginal:   ") 
	// fmt.Println(VideoList)
	// fmt.Print("After:   ") 
	// fmt.Println(PublishVideoList)
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: PublishVideoList,
	})
	//fmt.Printf("publish list\n")
}
