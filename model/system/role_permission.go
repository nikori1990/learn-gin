package system

import (
	"learn-gin/model/base"
)

type RolePermission struct {
	base.Model
	RoleId       string
	PermissionId string
}
