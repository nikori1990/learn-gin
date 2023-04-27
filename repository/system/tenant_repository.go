package system

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type TenantRepository struct {
}

func (TenantRepository) Create(tenant *system.Tenant) uint {
	if err := global.DB.Create(&tenant).Error; err != nil {
		panic(err)
	}
	return tenant.ID
}
