package controller

import (
	"GoDance/model"
	"net/http"
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

// CommentAction no practical effect, just check if token is valid
// CommentAction
// @Summary 评论操作
// @Description 登录用户对视频进行评论
// @Tags 互动接口
// @param token query string true "用户鉴权token"
// @param video_id query string true "视频id"
// @param action_type query string true "1-发布评论，2-删除评论"
// @param comment_text query string false "用户填写的评论内容，在action_type=1的时候使用"
// @param comment_id query string false "要删除的评论id，在action_type=2的时候使用"
// @Success 200 {string} success
// @Router /comment/action/ [post]
func CommentAction(c *gin.Context) {
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	if user, exist := usersLoginInfo[token]; exist {
		var comment model.Comment
		if actionType == "1" { //新增
			comment.UserId = user.Id
			comment.VideoId = int64(videoId)
			text := c.Query("comment_text")
			comment.CommentText = text
			comment.CreateDate = time.Now()
			model.InsertComment(&comment)
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0, StatusMsg: "评论成功"},
				Comment: Comment{
					Id:         comment.Id,
					User:       user,
					Content:    text,
					CreateDate: comment.CreateDate.Format("2006-01-02 15:04:05"),
				}})
			return
		} else if actionType == "2" { //删除
			commentId, _ := strconv.Atoi(c.Query("comment_id"))
			comment = model.FindCommentById(int64(commentId))
			if comment.UserId == user.Id { //是该评论主人
				model.DeleteComment(int64(commentId))
				c.JSON(http.StatusOK, Response{StatusCode: 0, StatusMsg: "评论删除成功"})
			} else { //不能删除他人评论
				c.JSON(http.StatusForbidden, Response{StatusCode: 1, StatusMsg: "没有权限删除他人评论"})
			}
		} else {
			c.JSON(http.StatusBadRequest, Response{StatusCode: 1, StatusMsg: "评论状态参数错误"})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// CommentList all videos have same demo comment list
// CommentList
// @Summary 评论列表
// @Description 查看视频的所有评论，按发布时间倒序
// @Tags 互动接口
// @param token query string true "用户鉴权token"
// @param video_id query string true "视频id"
// @Success 200 {string} success
// @Router /comment/list/ [get]
func CommentList(c *gin.Context) {
	token := c.Query("token")
	videoId, _ := strconv.Atoi(c.Query("video_id"))
	if user, exist := usersLoginInfo[token]; exist {
		comments := model.FindCommentsByVideoId(int64(videoId))
		var commentsVO = make([]Comment, len(comments))
		for i := 0; i < len(comments); i++ {
			comment := comments[i]
			commentsVO[i] = Comment{comment.Id, user, comment.CommentText, comment.CreateDate.Format("2006-01-02 15:04:05")}
		}
		c.JSON(http.StatusOK, CommentListResponse{
			Response:    Response{StatusCode: 0, StatusMsg: "获取评论列表成功"},
			CommentList: commentsVO,
		})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}
