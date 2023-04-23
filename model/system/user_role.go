package system

import (
	"learn-gin/model/base"
)

type UserRole struct {
	base.Model
	UserId   uint   `json:"userId"`
	RoleId   uint   `json:"roleId"`
	RoleName string `json:"roleName"`
	Username string `json:"username"`
}
