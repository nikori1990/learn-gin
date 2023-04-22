package dao

import (
	"learn-gin/global"
	"learn-gin/models/system"
)

type MenuDao struct {
}

func (dao MenuDao) GetById(id string) *system.Menu {
	var menu *system.Menu
	global.DB.Where("id=?", id).First(&menu)
	return menu
}
