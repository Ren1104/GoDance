package controller

import (
	"net/http"
	"simple-demo/config"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoID, err := strconv.Atoi(c.Query("video_id"))
	if err != nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "video does not exist"}) //video不存在
		return
	}
	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {

			text := c.Query("comment_text")
			const DateOnly = "2006-01-02"
			createDate := time.Now().Format(DateOnly)
			dbCom := config.Comment{VideoID: uint(videoID), UserID: uint(user.Id), Content: text}

			if err := config.CommentAdd(dbCom); err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "comment failed"}) //添加評論失敗
				return
			} else {
				c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
					Comment: Comment{
						Id:         int64(videoID),
						User:       user,
						Content:    text,
						CreateDate: createDate,
					}})
				return
			}

		} else if actionType == "2" {

			comId, err := strconv.Atoi(c.Query("comment_id"))
			if err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "comment does not exist"})
				return
			}

			if err := config.CommentDelete(comId); err != nil {
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "delete comment failed"}) //刪除失敗
				return
			} else {
				c.JSON(http.StatusOK, Response{StatusCode: 0}) //正確刪除
				return
			}

		} else {

			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "invaild actionType"}) //動作不支持
			return
		}
	} else {

		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"}) //token无效返回错误
		return
	}
}

// CommentList all videos have same demo comment list
// 视频列表接口
func CommentList(c *gin.Context) {
	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: DemoComments,
	})
}
