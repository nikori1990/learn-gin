package system

import (
	"learn-gin/enum"
	"learn-gin/model/base"
)

type Permission struct {
	base.Model
	Name     string              `json:"name"`
	Path     string              `json:"path"` // url /api/users  /api/users/:id
	Type     enum.PermissionType `json:"type"` // '权限类型，页面-1，按钮-2',
	Sort     int                 `json:"sort"`
	Method   string              `json:"method"` // GET DELETE
	ParentId uint                `json:"parentId"`
}
