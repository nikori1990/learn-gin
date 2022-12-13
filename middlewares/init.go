package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
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

// JWTAuth 鉴权中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		fmt.Println("requestUri=", uri)
		if uri == "/api/v1/login" {
			c.Next()
		}

		// 获取请求头中 token，实际是一个完整被签名过的 token；a complete, signed token
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == "" {
			c.JSON(http.StatusForbidden, "No token. You don't have permission!")
			c.Abort()
			return
		}

		// 解析拿到完整有效的 token，里头包含解析后的 3 segment
		token, err := ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusForbidden, "Invalid token! You don't have permission!")
			c.Abort()
			return
		}
		// 获取 token 中的 claims
		claims, ok := token.Claims.(*AuthClaims)
		if !ok {
			c.JSON(http.StatusForbidden, "Invalid token!")
			c.Abort()
			return
		}

		// 将 claims 中的用户信息存储在 context 中
		c.Set("userId", claims.UserId)

		// 这里执行路由 HandlerFunc
		c.Next()
	}
}

type AuthClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.RegisteredClaims
	UserId      string   `json:"user_id"`
	Permissions []string `json:"permissions"`
}

var secretKey = []byte("some string")

// ParseToken 解析请求头中的 token string，转换成被解析后的 jwt.Token
func ParseToken(tokenStr string) (*jwt.Token, error) {
	// 解析 token string 拿到 token jwt.Token
	return jwt.ParseWithClaims(tokenStr, &AuthClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

// GenerateToken 一般在登录之后使用来生成 token 能够返回给前端
func GenerateToken(userId string, expireTime time.Time) (string, error) {
	// 创建一个 claim
	claim := AuthClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			// 过期时间
			ExpiresAt: jwt.NewNumericDate(expireTime),
			// 签名时间
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// 签名颁发者
			Issuer: "abcnull",
			// 签名主题
			Subject: "gindemo",
		},
	}

	// 使用指定的签名加密方式创建 token，有 1，2 段内容，第 3 段内容没有加上
	noSignedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// 使用 secretKey 密钥进行加密处理后拿到最终 token string
	return noSignedToken.SignedString(secretKey)
}
