package system

import (
	"github.com/gin-gonic/gin"
	"learn-gin/global"
	"learn-gin/model/api"
	"learn-gin/model/system"
	"learn-gin/model/system/query"
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
	global.Logger.Infof("add permission: %v success", permission)
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
	var permission system.Permission

	if err := c.ShouldBindJSON(&permission); err != nil {
		panic(err)
	}

	deleteId := permissionRepository.Delete(permission.ID)
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

func (PermissionService) Tree(c *gin.Context) {
	list := permissionRepository.List()
	tree := recursivePermission(list, 0)
	api.Success(c, tree)
}

func (PermissionService) Page(c *gin.Context) {
	var pageQuery query.PermissionQuery

	if err := c.ShouldBindQuery(&pageQuery); err != nil {
		panic(err)
	}

	pageQuery.SetFromLimit()
	page := permissionRepository.Page(&pageQuery)

	api.Success(c, page)
}

func recursivePermission(list []*system.Permission, pId uint) []*system.Permission {
	res := make([]*system.Permission, 0)

	for _, item := range list {
		if item.ParentId == pId {
			item.Children = recursivePermission(list, item.ID)
			if len(item.Children) == 0 {
				item.Children = nil
			}
			res = append(res, item)
		}
	}
	return res
}
