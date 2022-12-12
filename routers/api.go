package routers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/routers/apis"
	"net/http"
)

func ApiRoutersInit(router *gin.Engine) {
	apiRouter := router.Group("/api")
	{
		apiRouter.GET("/", func(context *gin.Context) {
			context.String(http.StatusOK, "我是一个api接口")
		})

		apis.MenuRoutersInit(apiRouter)
		apis.UserRoutersInit(apiRouter)
		apis.ArticleRoutersInit(apiRouter)
		apis.LoginRoutersInit(apiRouter)
	}
}
