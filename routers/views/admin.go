package views

import (
	"github.com/gin-gonic/gin"
	"learn-gin/routers/views/admin"
	"net/http"
)

func AdminRoutersInit(router *gin.RouterGroup) {
	viewsAdminRouter := router.Group("/admin")
	{
		viewsAdminRouter.GET("", func(context *gin.Context) {
			context.HTML(http.StatusOK, "admin/index.tmpl", gin.H{
				"title": "首页",
			})
		})

		admin.ArticleRoutersInit(viewsAdminRouter)
		admin.NewsRoutersInit(viewsAdminRouter)
	}
}
