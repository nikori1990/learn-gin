package router

import (
	"github.com/gin-gonic/gin"
)

type RoleRouter struct {
}

func (RoleRouter) Init(group *gin.RouterGroup) {
	roleRouter := group.Group("/roles")
	{
		roleRouter.GET("", roleService.List)
		roleRouter.GET("/:id", roleService.GetById)
		roleRouter.POST("", roleService.Create)
		roleRouter.PUT("", roleService.Update)
		roleRouter.DELETE("/:id", roleService.Delete)
		roleRouter.GET("/:id/permissions", roleService.ListPermissions)
	}
}
