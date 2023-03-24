package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service/v1"
)

type MenuRouter struct {
}

func (router *MenuRouter) Init(group *gin.RouterGroup) {
	menuRouter := group.Group("/menus")
	{
		menuRouter.GET("", v1.MenuService{}.List)
		menuRouter.GET("/:id", v1.MenuService{}.Get)
	}
}
