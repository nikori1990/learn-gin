package base

import "github.com/golang-jwt/jwt/v4"

type AuthClaims struct { // token里面添加用户信息，验证token后可能会用到用户信息
	jwt.RegisteredClaims
	TenantId    string     `json:"tenantId"`
	Username    string     `json:"username"`
	Role        string     `json:"role"`
	Permissions [][]string `json:"permissions"`
}
