package system

import (
	"learn-gin/enum"
	"learn-gin/model/base"
)

type Menu struct {
	base.Model
	Name      string    `json:"name"`
	Path      string    `json:"path"` // /dashboard /system/users
	Icon      string    `json:"icon"`
	Sort      int       `json:"sort"`
	KeepAlive enum.Flag `json:"keepAlive"`
	ParentId  uint      `json:"parentId"`
	Children  []*Menu   `json:"children" gorm:"-"`
	TenantId  uint      `json:"tenantId"`
}
