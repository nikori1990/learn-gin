package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"learn-gin/global"
	"learn-gin/model/api"
	"learn-gin/model/base"
	"time"
)

func GetClaims(c *gin.Context) (*base.AuthClaims, error) {
	// 获取请求头中 token，实际是一个完整被签名过的 token；a complete, signed token
	authorization := c.GetHeader("Authorization")
	if authorization == "" {
		api.Forbidden(c, "No token. You don't have permission!")
		c.Abort()
	}

	tokenRune := []rune(authorization)
	//fmt.Println("tokenRune", string(tokenRune[7:]))
	//fmt.Println("tokenRune", string(tokenRune[:7]))

	// 解析拿到完整有效的 token，里头包含解析后的 3 segment
	token, err := ParseToken(string(tokenRune[7:]))
	if err != nil {
		api.Forbidden(c, "Invalid token! You don't have permission!")
		c.Abort()
	}
	// 获取 token 中的 claims
	claims, ok := token.Claims.(*base.AuthClaims)
	if !ok {
		api.Forbidden(c, "Invalid token!")
		c.Abort()
	}
	return claims, nil
}

// ParseToken 解析请求头中的 token string，转换成被解析后的 jwt.Token
func ParseToken(tokenStr string) (*jwt.Token, error) {
	// 解析 token string 拿到 token jwt.Token
	secretKey := []byte(global.CONFIG.Security.JWT.SecretKey)
	fmt.Println("secretKey:", secretKey)
	return jwt.ParseWithClaims(tokenStr, &base.AuthClaims{}, func(tk *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
}

// GenerateToken 一般在登录之后使用来生成 token 能够返回给前端
func GenerateToken(tenantId string, username string, permissions [][]string) (string, error) {

	now := time.Now()
	expireTime := now.Add(time.Hour)

	secretKey := []byte(global.CONFIG.Security.JWT.SecretKey)
	// 创建一个 claim
	claim := base.AuthClaims{
		TenantId:    tenantId,
		Username:    username,
		Permissions: permissions,
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
