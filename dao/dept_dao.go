package dao

import (
	"learn-gin/global"
	"learn-gin/models/system"
)

type DeptDao struct {
}

func (dao DeptDao) Create(dept *system.Dept) uint {
	if err := global.DB.Create(&dept).Error; err != nil {
		panic(err)
	}
	return dept.ID
}
