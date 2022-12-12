package views

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GoodsRoutersInit(router *gin.RouterGroup) {
	goodsRouter := router.Group("/goods")
	{
		goodsRouter.GET("", func(context *gin.Context) {
			context.HTML(http.StatusOK, "default/goods.tmpl", gin.H{
				"title": "我是商品页面",
				"price": 20,
			})
		})
	}
}
