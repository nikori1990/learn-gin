package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/model/api"
	"learn-gin/model/system"
)

type RolePermissionService struct {
}

func (RolePermissionService) Create(c *gin.Context) {
	var rolePermission system.RolePermission
	if err := c.ShouldBindJSON(&rolePermission); err != nil {
		panic(err)
	}
	createId := rolePermissionRepository.Create(&rolePermission)
	api.Success(c, createId)
}

func (RolePermissionService) List(c *gin.Context) {
	list := rolePermissionRepository.List()
	api.Success(c, list)
}
