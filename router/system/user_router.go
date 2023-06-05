package system

import (
	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserRouter struct {
}

func (UserRouter) Init(group *gin.RouterGroup) {
	router := group.Group("/users")
	{
		router.GET("", userService.List)
		router.GET("/:id", userService.GetById)
		router.POST("", userService.Create)
		router.PUT("", userService.Update)
		router.DELETE("", userService.Delete)
		router.GET("/page", userService.Page)
	}
}
