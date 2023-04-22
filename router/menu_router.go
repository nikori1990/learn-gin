package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service"
)

type MenuRouter struct {
}

var menuService = new(service.MenuService)

func (router *MenuRouter) Init(group *gin.RouterGroup) {
	menuRouter := group.Group("/menus")
	{
		menuRouter.GET("", menuService.List)
		menuRouter.GET("/:id", menuService.Get)
	}
}
