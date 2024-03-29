package system

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type UserRoleRepository struct {
}

func (UserRoleRepository) Create(userRole *system.UserRole) uint {
	if err := global.DB.Create(&userRole).Error; err != nil {
		panic(err)
	}
	return userRole.ID
}

func (UserRoleRepository) List() []*system.UserRole {
	var list []*system.UserRole
	if err := global.DB.Find(&list).Error; err != nil {
		panic(err)
	}
	return list
}

func (UserRoleRepository) ListByUserId(id uint) []*system.UserRole {
	var list []*system.UserRole
	if err := global.DB.Where("user_id = ?", id).Find(&list).Error; err != nil {
		panic(err)
	}
	return list
}
