package system

import (
	"learn-gin/model/base"
)

type Menu struct {
	base.Model
	Name     string  `json:"name"`
	Path     string  `json:"path"` // /dashboard /system/users
	Icon     string  `json:"icon"`
	ParentId uint    `json:"parentId"`
	Children []*Menu `json:"children" gorm:"-"`
}
