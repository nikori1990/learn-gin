package system

import (
	"github.com/gin-gonic/gin"
)

type PermissionRouter struct {
}

func (PermissionRouter) Init(group *gin.RouterGroup) {
	router := group.Group("/permissions")
	{
		router.GET("", permissionService.List)
		router.GET("/:id", permissionService.GetById)
		router.POST("", permissionService.Create)
		router.PUT("", permissionService.Update)
		router.DELETE("", permissionService.Delete)
		router.GET("/page", permissionService.Page)
		router.GET("/tree", permissionService.Tree)
	}
}
