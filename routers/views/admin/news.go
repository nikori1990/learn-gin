package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NewsRoutersInit(group *gin.RouterGroup) {
	newsRouter := group.Group("/news")
	{
		newsRouter.GET("", func(context *gin.Context) {
			context.HTML(http.StatusOK, "admin/news.tmpl", gin.H{
				"title": "首页",
			})
		})
	}
}
