package system

import (
	"learn-gin/model/base"
)

type User struct {
	base.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Age      uint   `json:"age"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	TenantId uint   `json:"tenantId"`
}
