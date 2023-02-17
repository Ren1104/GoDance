package main

import (
	"GoDance/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(r *gin.Engine) {
	r.Static("/static", "./public")
	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	//apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.StaticFS("/static", http.Dir("./public"))

}
