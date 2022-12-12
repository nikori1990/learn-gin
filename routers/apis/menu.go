package apis

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/api"
)

func MenuRoutersInit(group *gin.RouterGroup) {
	menuRouter := group.Group("/menus")
	{
		menuRouter.GET("", api.MenuController{}.List)
		menuRouter.GET("/:id", api.MenuController{}.Get)
	}
}
