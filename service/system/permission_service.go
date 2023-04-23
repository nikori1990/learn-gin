package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
	"strconv"
)

type PermissionService struct {
}

func (PermissionService) Create(c *gin.Context) {
	var permission system.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		panic(err)
	}

	permissionId := permissionRepository.Create(&permission)
	api.Success(c, permissionId)
}

func (PermissionService) Update(c *gin.Context) {
	var permission system.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		panic(err)
	}

	updateId := permissionRepository.Update(&permission)
	api.Success(c, updateId)
}

func (PermissionService) Delete(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	deleteId := permissionRepository.Delete(uint(id))
	api.Success(c, deleteId)
}

func (PermissionService) GetById(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		panic(err)
	}
	permission := permissionRepository.GetById(uint(id))
	api.Success(c, permission)
}

func (PermissionService) List(c *gin.Context) {
	list := permissionRepository.List()
	api.Success(c, list)
}
