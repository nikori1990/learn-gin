package system

import (
	"github.com/gin-gonic/gin"
)

type PermissionRouter struct {
}

func (PermissionRouter) Init(group *gin.RouterGroup) {
	permissionRouter := group.Group("/permissions")
	{
		permissionRouter.GET("", permissionService.List)
		permissionRouter.GET("/:id", permissionService.GetById)
		permissionRouter.POST("", permissionService.Create)
		permissionRouter.PUT("", permissionService.Update)
		permissionRouter.DELETE("", permissionService.Delete)
		permissionRouter.GET("/page", permissionService.Page)
		permissionRouter.GET("/tree", permissionService.Tree)
	}
}
