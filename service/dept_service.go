package service

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
	"learn-gin/repository"
	"strconv"
)

type DeptService struct {
}

var deptRepository repository.DeptRepository

func (service DeptService) Create(c *gin.Context) {
	var dept system.Dept

	if err := c.ShouldBindJSON(&dept); err != nil {
		panic(err)
	}

	deptId := deptRepository.Create(&dept)
	api.Success(c, deptId)
}

func (service DeptService) Update(c *gin.Context) {
	var dept system.Dept

	if err := c.ShouldBindJSON(&dept); err != nil {
		panic(err)
	}

	updateId := deptRepository.Update(&dept)
	api.Success(c, updateId)
}

func (service DeptService) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	deleteId := deptRepository.Delete(uint(id))
	api.Success(c, deleteId)
}

func (service DeptService) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	dept := deptRepository.GetById(uint(id))
	api.Success(c, dept)
}

func (service DeptService) List(c *gin.Context) {
	list := deptRepository.List()
	api.Success(c, list)
}
