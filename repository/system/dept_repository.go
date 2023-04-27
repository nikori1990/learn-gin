package system

import (
	"learn-gin/global"
	"learn-gin/model/system"
)

type DeptRepository struct {
}

func (DeptRepository) Create(dept *system.Dept) uint {
	if err := global.DB.Create(&dept).Error; err != nil {
		panic(err)
	}
	return dept.ID
}

func (DeptRepository) Update(dept *system.Dept) uint {
	if err := global.DB.Model(&system.Dept{}).Where("id = ?", dept.ID).Update("name", dept.Name).Error; err != nil {
		panic(err)
	}
	return dept.ID
}

func (DeptRepository) Delete(id uint) uint {
	if err := global.DB.Delete(&system.Dept{}, id).Error; err != nil {
		panic(err)
	}
	return id
}

func (DeptRepository) GetById(id uint) *system.Dept {
	var dept *system.Dept
	if err := global.DB.Limit(1).Find(&dept, id).Error; err != nil {
		panic(err)
	}
	return dept
}

func (DeptRepository) List() []*system.Dept {
	var deptList []*system.Dept
	if err := global.DB.Find(&deptList).Error; err != nil {
		panic(err)
	}
	return deptList
}
