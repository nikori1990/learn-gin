package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
	"strconv"
)

type DeptService struct {
}

func (DeptService) Create(c *gin.Context) {
	var dept system.Dept

	if err := c.ShouldBindJSON(&dept); err != nil {
		panic(err)
	}

	deptId := deptRepository.Create(&dept)
	api.Success(c, deptId)
}

func (DeptService) Update(c *gin.Context) {
	var dept system.Dept

	if err := c.ShouldBindJSON(&dept); err != nil {
		panic(err)
	}

	updateId := deptRepository.Update(&dept)
	api.Success(c, updateId)
}

func (DeptService) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	deleteId := deptRepository.Delete(uint(id))
	api.Success(c, deleteId)
}

func (DeptService) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	dept := deptRepository.GetById(uint(id))
	api.Success(c, dept)
}

func (DeptService) List(c *gin.Context) {
	list := deptRepository.List()
	api.Success(c, list)
}

func (DeptService) Page(c *gin.Context) {

}

func (DeptService) Tree(c *gin.Context) {
	list := deptRepository.List()
	tree := recursiveDept(list, 0)
	api.Success(c, tree)
}

func recursiveDept(list []*system.Dept, pId uint) []*system.Dept {
	res := make([]*system.Dept, 0)

	for _, item := range list {
		if item.ParentId == pId {
			item.Children = recursiveDept(list, item.ID)
			if len(item.Children) == 0 {
				item.Children = nil
			}
			res = append(res, item)
		}
	}
	return res
}
