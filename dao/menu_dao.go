package dao

import (
	"learn-gin/global"
	"learn-gin/models/system"
)

type MenuDao struct {
}

func (dao MenuDao) GetById(id string) *system.SysMenu {
	var menu *system.SysMenu
	global.DB.Where("id=?", id).First(&menu)
	return menu
}
