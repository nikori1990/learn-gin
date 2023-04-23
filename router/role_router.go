package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service"
)

type RoleRouter struct {
}

var roleService = new(service.RoleService)

func (RoleRouter) Init(group *gin.RouterGroup) {
	roleRouter := group.Group("/roles")
	{
		roleRouter.GET("", roleService.List)
		roleRouter.GET("/:id", roleService.GetById)
		roleRouter.POST("", roleService.Create)
		roleRouter.PUT("", roleService.Update)
		roleRouter.DELETE("/:id", roleService.Delete)
	}
}
