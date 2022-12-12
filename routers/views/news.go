package views

import (
	"github.com/gin-gonic/gin"
	"learn-gin/models"
	"net/http"
)

func NewsRoutersInit(router *gin.RouterGroup) {
	newRouter := router.Group("/news")
	{
		newRouter.GET("", func(context *gin.Context) {
			news := &models.Article{
				Title:   "新闻标题",
				Content: "新闻详情",
			}
			context.HTML(http.StatusOK, "default/news.tmpl", gin.H{
				"title": "新闻页面",
				"news":  news,
			})
		})
	}
}
