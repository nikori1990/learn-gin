package service

import (
	"learn-gin/dao"
	"learn-gin/models/system"
)

type DeptService struct {
}

var deptDao dao.DeptDao

func (service DeptService) save(dept *system.Dept) {
	deptDao.Create(dept)
}
