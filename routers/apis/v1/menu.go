package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/api/v1"
)

func MenuRoutersInit(group *gin.RouterGroup) {
	menuRouter := group.Group("/menus")
	{
		menuRouter.GET("", v1.MenuController{}.List)
		menuRouter.GET("/:id", v1.MenuController{}.Get)
	}
}
