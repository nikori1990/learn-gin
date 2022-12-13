package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/admin"
	"net/http"
)

func ArticleRoutersInit(group *gin.RouterGroup) {
	apiArticleRouter := group.Group("/articles")
	{
		apiArticleRouter.POST("/add", func(context *gin.Context) {
			title := context.PostForm("title")
			desc := context.PostForm("desc")

			fmt.Println(title)
			fmt.Println(desc)

			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"message": "成功",
			})
		})

		apiArticleRouter.GET("/edit", admin.NewsController{}.Edit)
	}
}
