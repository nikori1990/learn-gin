package system

import (
	"learn-gin/model/base"
)

type RolePermission struct {
	base.Model
	RoleId       uint `json:"roleId"`
	PermissionId uint `json:"permissionId"`
}
