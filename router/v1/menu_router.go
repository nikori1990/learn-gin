package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service/v1"
)

type MenuRouter struct {
}

var menuService = new(v1.MenuService)

func (router *MenuRouter) Init(group *gin.RouterGroup) {
	menuRouter := group.Group("/menus")
	{
		menuRouter.GET("", menuService.List)
		menuRouter.GET("/:id", menuService.Get)
	}
}
