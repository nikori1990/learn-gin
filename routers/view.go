package routers

import (
	"github.com/gin-gonic/gin"
	"learn-gin/middlewares"
	"learn-gin/routers/views"
)

func ViewRoutersInit(router *gin.Engine) {
	viewRouter := router.Group("/", middlewares.InitMiddleware)
	{
		views.AdminRoutersInit(viewRouter)
		views.GoodsRoutersInit(viewRouter)
		views.NewsRoutersInit(viewRouter)
	}
}
