package system

import (
	"github.com/gin-gonic/gin"
)

type RolePermissionRouter struct {
}

func (RolePermissionRouter) Init(router *gin.RouterGroup) {
	rolePermissionRouter := router.Group("/rolePermissions")
	{
		rolePermissionRouter.GET("", rolePermissionService.List)
		rolePermissionRouter.POST("", rolePermissionService.Create)
	}
}
