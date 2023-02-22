package main

import (
	"GoDance/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func initRouter(r *gin.Engine) {

	apiRouter := r.Group("/douyin")

	apiRouter.GET("/user/", controller.UserInfo)
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.GET("/publish/list/", controller.PublishList)
	apiRouter.GET("/comment/list/", controller.CommentList)
	apiRouter.POST("/user/register/", controller.Register)
	apiRouter.POST("/user/login/", controller.Login)
	apiRouter.POST("/publish/action/", controller.Publish)
	apiRouter.POST("/comment/action/", controller.CommentAction)
	apiRouter.POST("/favorite/action/", controller.FavoriteAction)
	apiRouter.StaticFS("/public", http.Dir("./public"))
}
