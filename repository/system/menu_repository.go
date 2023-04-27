package system

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type MenuRepository struct {
}

func (MenuRepository) GetById(id string) *system.Menu {
	var menu *system.Menu
	global.DB.Where("id=?", id).First(&menu)
	return menu
}
