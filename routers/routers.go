package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"learn-gin/middlewares"
	"learn-gin/models"
)

func globalMiddleware(c *gin.Context) {
	fmt.Println("这是全局中间件")
	//request := c.Request
	//fmt.Println(request)
}

func SetUpRouters() *gin.Engine {
	router := gin.Default()
	// 自定义模板函数
	router.SetFuncMap(template.FuncMap{
		"UnixToTime": models.UnixToTime,
		"ConcatStr":  models.ConcatStr,
	})
	// 模板
	router.LoadHTMLGlob("templates/**/*")
	// 静态web
	router.Static("/static", "./static")

	router.Use(globalMiddleware)
	router.Use(middlewares.JWTAuth())

	//router.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "值: %v", "你好, gin")
	//})
	//router.GET("/news", func(context *gin.Context) {
	//	context.String(http.StatusOK, "我是新闻页面1")
	//})

	DefaultRoutersInit(router)
	ViewRoutersInit(router)
	ApiRoutersInit(router)

	return router
}
