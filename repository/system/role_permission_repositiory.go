package system

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type RolePermissionRepository struct {
}

func (RolePermissionRepository) Create(rolePermission *system.RolePermission) uint {
	if err := global.DB.Create(&rolePermission).Error; err != nil {
		panic(err)
	}
	return rolePermission.ID
}

func (RolePermissionRepository) Update(permission *system.RolePermission) uint {
	return 0
}

func (RolePermissionRepository) List() []*system.RolePermission {
	var list []*system.RolePermission
	if err := global.DB.Find(&list).Error; err != nil {
		panic(err)
	}
	return list
}

func (RolePermissionRepository) ListByRoleId(roleId uint) []*system.RolePermission {
	var list []*system.RolePermission
	if err := global.DB.Find(&list).Where("role_id = ?", roleId).Error; err != nil {
		panic(err)
	}
	return list
}

func (RolePermissionRepository) ListByRoleIds(ids []uint) []*system.RolePermission {
	var list []*system.RolePermission
	if err := global.DB.Find(&list).Where("role_id IN ?", ids).Error; err != nil {
		panic(err)
	}
	return list
}
