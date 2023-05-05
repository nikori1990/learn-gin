package router

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service"
)

type LoginRouter struct {
}

var loginService = new(service.LoginService)

func (router *LoginRouter) Init(group *gin.RouterGroup) {
	loginRouter := group.Group("/login")
	{
		loginRouter.POST("/login", loginService.Login)
		loginRouter.POST("/logout", loginService.Logout)
	}
}
