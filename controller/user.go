package controller

import (
	"GoDance/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync/atomic"
	_ "sync/atomic"
)

var userIdSequence = int64(1)

type UserResponse struct {
	Response
	User model.User `json:"user"`
}

type LoginResponse struct {
	Response
	Id    int64  `json:"user_id,omitempty"`
	Token string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	token := username + password
	var user model.User
	db.Find(&user, "username=?", username)
	if user.Id != 0 {
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
		})
	} else {
		atomic.AddInt64(&userIdSequence, 1)
		user := model.User{
			Id:       userIdSequence,
			Username: username,
			Password: password,
			Nickname: username,
			Token:    token,
		}
		fmt.Println(user)
		if db.Create(&user).RowsAffected == 1 {
			c.JSON(http.StatusOK, LoginResponse{
				Response: Response{StatusCode: 0},
				Id:       userIdSequence,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, LoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "Something wrong"},
			})
		}
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	var user model.User
	db.Where("username=? AND password=?", username, password).Find(&user)
	if user.Id != 0 {
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Login success!"},
			Id:       user.Id,
			Token:    user.Token,
		})
	} else {
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	var user model.User
	db.Find(&user, "token=?", token)
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     user,
	})
}
