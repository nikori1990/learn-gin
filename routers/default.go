package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/models"
	"net/http"
	"time"
)

func initMiddleware(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1 --- 我是一个中间件")
	c.Next()
	fmt.Println("2 --- 我是一个中间件")
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func initMiddleware2(c *gin.Context) {
	start := time.Now().UnixNano()
	fmt.Println("1 --- 我是一个中间件2")
	c.Next()
	fmt.Println("2 --- 我是一个中间件2")
	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func DefaultRoutersInit(router *gin.Engine) {

	defaultRouter := router.Group("/")
	{
		defaultRouter.GET("", initMiddleware, initMiddleware2, func(context *gin.Context) {
			fmt.Println("我是一个首页")
			context.HTML(http.StatusOK, "default/index.tmpl", gin.H{
				"title":   "首页",
				"message": "描述",
				"hobby":   []string{"吃饭", "睡觉", "写代码"},
				"date":    1647106181,
			})
		})

		// curl -X POST
		defaultRouter.POST("/add", func(context *gin.Context) {
			context.String(http.StatusOK, "我是POST请求")
		})
		defaultRouter.PUT("/edit", func(context *gin.Context) {
			context.String(http.StatusOK, "这是一个put请求")
		})
		defaultRouter.DELETE("/delete", func(context *gin.Context) {
			context.String(http.StatusOK, "这是一个delete请求")
		})

		defaultRouter.GET("/json", func(context *gin.Context) {
			context.JSON(http.StatusOK, map[string]interface{}{
				"success": true,
				"msg":     "你好gin",
			})
		})

		defaultRouter.GET("/json2", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"success": true,
				"msg":     "你好 gin -- 22",
			})
		})

		defaultRouter.GET("/json3", func(context *gin.Context) {
			a := &models.Article{
				Title:   "我是一个标题",
				Desc:    "描述",
				Content: "测试内容",
			}

			context.JSON(http.StatusOK, a)
		})

		defaultRouter.GET("/jsonp", func(context *gin.Context) {
			a := &models.Article{
				Title:   "我是一个标题 -- jsonp",
				Desc:    "描述",
				Content: "测试内容",
			}

			context.JSONP(http.StatusOK, a)
		})

		defaultRouter.GET("/xml", func(context *gin.Context) {
			context.XML(http.StatusOK, gin.H{
				"success": true,
				"msg":     "你好 gin -- 22",
			})
		})
	}
}
