package service

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
	"learn-gin/repository"
	"strconv"
)

type RoleService struct {
}

var roleRepository = new(repository.RoleRepository)

func (RoleService) Create(c *gin.Context) {
	var role system.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		panic(err)
	}
	createId := roleRepository.Create(&role)
	api.Success(c, createId)
}

func (RoleService) Update(c *gin.Context) {
	var role system.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		panic(err)
	}
	updateId := roleRepository.Update(&role)
	api.Success(c, updateId)
}

func (RoleService) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	deleteId := roleRepository.Delete(uint(id))
	api.Success(c, deleteId)
}

func (RoleService) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	role := roleRepository.GetById(uint(id))
	api.Success(c, role)
}

func (RoleService) List(c *gin.Context) {
	list := roleRepository.List()
	api.Success(c, list)
}
