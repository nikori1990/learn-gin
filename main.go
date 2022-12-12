package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"html/template"
	"learn-gin/models"
	"learn-gin/routers"
)

func globalMiddleware(c *gin.Context) {
	fmt.Println("这是全局中间件")
	//request := c.Request
	//fmt.Println(request)
}

func main() {
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

	//router.GET("/", func(context *gin.Context) {
	//	context.String(http.StatusOK, "值: %v", "你好, gin")
	//})
	//router.GET("/news", func(context *gin.Context) {
	//	context.String(http.StatusOK, "我是新闻页面1")
	//})

	routers.DefaultRoutersInit(router)
	routers.ViewRoutersInit(router)
	routers.ApiRoutersInit(router)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	configErr := viper.ReadInConfig()
	if configErr != nil {
		panic(fmt.Errorf("fatal error config file: %w", configErr))
	}

	fmt.Println("name", viper.Get("name"))

	routerErr := router.Run(":8000")
	if routerErr != nil {
		return
	}
}
