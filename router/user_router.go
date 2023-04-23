package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service"
)

type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type UserRouter struct {
}

var userService = new(service.UserService)

func (router *UserRouter) Init(group *gin.RouterGroup) {
	userRouter := group.Group("/users")
	{
		userRouter.GET("", userService.List)
		userRouter.GET("/:id", userService.GetById)
		userRouter.POST("", userService.Create)
		userRouter.PUT("", userService.Update)
		userRouter.DELETE("", userService.Delete)
	}
}
