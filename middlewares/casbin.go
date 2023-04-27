package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/model/api"
	"learn-gin/utils"
)

func CasbinHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		fmt.Println("requestUri=", uri)

		ignoreUris := global.CONFIG.Security.IgnoreUris
		fmt.Println("ignoreUris=", ignoreUris)

		if utils.Contains(ignoreUris, uri) {
			c.Next()
			return
		}

		claims, _ := utils.GetClaims(c)
		// 获取用户
		sub := claims.Username
		//获取请求的PATH
		obj := c.Request.URL.Path
		// 获取请求方法
		act := c.Request.Method
		//
		domain := claims.TenantId

		success, _ := global.CasbinEnforcer.Enforce(sub, domain, obj, act)
		if !success {
			api.Forbidden(c, "权限不足")
			c.Abort()
			return
		}

		c.Next()
	}
}
