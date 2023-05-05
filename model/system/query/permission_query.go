package query

import (
	"learn-gin/enum"
	"learn-gin/model/base"
)

type PermissionQuery struct {
	base.PageQuery
	Name string              `json:"name" form:"name"`
	Path string              `json:"path" form:"path"`
	Type enum.PermissionType `json:"type" form:"type"`
}
