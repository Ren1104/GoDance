package main

import (
	"GoDance/controller"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
}
