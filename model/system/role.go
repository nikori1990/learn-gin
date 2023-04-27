package system

import "learn-gin/model/base"

type Role struct {
	base.Model
	Name     string `json:"name"` // admin operator guest
	Code     string `json:"code"`
	TenantId uint   `json:"tenantId"`
}
