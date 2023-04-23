package router

import (
	"github.com/gin-gonic/gin"
)

type UserRoleRouter struct {
}

func (UserRoleRouter) Init(router *gin.RouterGroup) {
	userRoleRouter := router.Group("/userRoles")
	{
		userRoleRouter.GET("", userRoleService.List)
		userRoleRouter.POST("", userRoleService.Create)
	}
}
