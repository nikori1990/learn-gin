package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/utils"
)

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
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

		global.Logger.Infof("username: %s", claims.Username)
		global.Logger.Infof("permissions: %s", claims.Permissions)

		// 将 claims 中的用户信息存储在 context 中
		c.Set("username", claims.Username)
		c.Set("permissions", claims.Permissions)

		// 这里执行路由 HandlerFunc
		c.Next()
	}
}
