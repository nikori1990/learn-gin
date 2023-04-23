package router

import (
	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
}

func (router *MenuRouter) Init(group *gin.RouterGroup) {
	menuRouter := group.Group("/menus")
	{
		menuRouter.GET("", menuService.List)
		menuRouter.GET("/:id", menuService.Get)
	}
}
