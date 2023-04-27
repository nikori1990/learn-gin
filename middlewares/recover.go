package middlewares

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"log"
	"runtime/debug"
)

func Recover() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				//打印错误堆栈信息
				log.Printf("panic: %v\n", r)
				debug.PrintStack()
				//封装通用json返回
				//c.JSON(http.StatusOK, Result.Fail(errorToString(r)))
				//Result.Fail不是本例的重点，因此用下面代码代替
				api.Failed(c, errorToString(r))
				//终止后续接口调用，不加的话recover到异常后，还会继续执行接口里后续代码
				c.Abort()
			}
		}()
		//加载完 defer recover，继续后续接口调用
		c.Next()
	}
}

func errorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
