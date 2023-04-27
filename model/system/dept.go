package system

import (
	"learn-gin/model/base"
)

type Dept struct {
	base.Model
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	ParentId uint   `json:"parentId"`
	TenantId uint   `json:"tenantId"`
}
