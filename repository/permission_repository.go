package repository

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type PermissionRepository struct {
}

func (PermissionRepository) Create(permission *system.Permission) uint {
	if err := global.DB.Create(&permission).Error; err != nil {
		panic(err)
	}
	return permission.ID
}

func (PermissionRepository) Update(permission *system.Permission) uint {
	if err := global.DB.Model(&system.Permission{}).Where("id = ?", permission.ID).Updates(system.Permission{Path: permission.Path, Method: permission.Method}).Error; err != nil {
		panic(err)
	}
	return permission.ID
}

func (PermissionRepository) Delete(id uint) uint {
	if err := global.DB.Delete(&system.Permission{}, id).Error; err != nil {
		panic(err)
	}
	return id
}

func (PermissionRepository) GetById(id uint) *system.Permission {
	var permission *system.Permission
	if err := global.DB.Limit(1).Find(&permission, id).Error; err != nil {
		panic(err)
	}
	return permission
}

func (PermissionRepository) List() []*system.Permission {
	var permissionList []*system.Permission
	if err := global.DB.Find(&permissionList).Error; err != nil {
		panic(err)
	}
	return permissionList
}

func (PermissionRepository) ListByIds(ids []uint) []*system.Permission {
	var list []*system.Permission
	if err := global.DB.Find(&list, ids).Error; err != nil {
		panic(err)
	}
	return list
}
