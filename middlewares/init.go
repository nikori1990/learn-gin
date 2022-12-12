package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func InitMiddleware(c *gin.Context) {
	// 判断用户是否登陆
	fmt.Println(time.Now())

	fmt.Println(c.Request.URL)

	c.Set("username", "哈哈")

	// 定义一个goroutine 统计日志

	cCp := c.Copy()
	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path " + cCp.Request.URL.Path)
	}()
}
