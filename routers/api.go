package routers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/routers/apis/v1"
)

func ApiRoutersInit(router *gin.Engine) {
	//ginAccounts := gin.Accounts{
	//	"niko":  "123456",
	//	"admin": "Admin123",
	//}

	apiRouter := router.Group("/api/v1")
	{
		//apiRouter.GET("/", gin.BasicAuth(ginAccounts), func(context *gin.Context) {
		//	context.String(http.StatusOK, "我是一个api接口")
		//})

		v1.MenuRoutersInit(apiRouter)
		v1.UserRoutersInit(apiRouter)
		v1.ArticleRoutersInit(apiRouter)
		v1.LoginRoutersInit(apiRouter)
	}
}
