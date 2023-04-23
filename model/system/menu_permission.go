package system

import (
	"learn-gin/model/base"
)

type MenuPermission struct {
	base.Model
	MenuId       string
	PermissionId string
}
