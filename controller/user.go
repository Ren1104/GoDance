package controller

import (
	"GoDance/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.Register(username, password)
	if err != nil {
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Register success"},
			Id:       user.Id,
			Token:    user.Token,
		})
	}
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user, err := service.Login(username, password)
	if err == nil {
		fmt.Printf("success1")
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Login success!"},
			Id:       user.Id,
			Token:    user.Token,
		})
	} else {
		c.JSON(http.StatusOK, LoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	}
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	user, err := service.FindUserByToken(token)
	if err == nil {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 0},
			User:     user,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: err.Error()},
			User:     user,
		})
	}
}
