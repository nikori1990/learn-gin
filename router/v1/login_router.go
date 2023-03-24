package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service/v1"
)

type LoginRouter struct {
}

func (router *LoginRouter) Init(group *gin.RouterGroup) {
	loginRouter := group.Group("/login")
	{
		loginRouter.POST("", v1.LoginService{}.Login)
	}
}
