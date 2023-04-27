package system

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type RoleRepository struct {
}

func (RoleRepository) Create(role *system.Role) uint {
	if err := global.DB.Create(&role).Error; err != nil {
		panic(err)
	}
	return role.ID
}

func (RoleRepository) Update(role *system.Role) uint {
	if err := global.DB.Model(&system.Role{}).Where("id = ?", role.ID).Update("name", role.Name).Error; err != nil {
		panic(err)
	}
	return role.ID
}

func (RoleRepository) Delete(id uint) uint {
	if err := global.DB.Delete(&system.Role{}, id).Error; err != nil {
		panic(err)
	}
	return id
}

func (RoleRepository) GetById(id uint) *system.Role {
	var role *system.Role
	if err := global.DB.First(&role, id).Error; err != nil {
		panic(err)
	}
	return role
}

func (RoleRepository) List() []*system.Role {
	var roles []*system.Role
	if err := global.DB.Find(&roles).Error; err != nil {
		panic(err)
	}
	return roles
}

func (RoleRepository) ListByIds(ids []uint) []*system.Role {
	var roles []*system.Role
	if err := global.DB.Find(&roles, ids).Error; err != nil {
		panic(err)
	}
	return roles
}
