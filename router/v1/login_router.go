package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/service/v1"
)

func LoginRoutersInit(group *gin.RouterGroup) {
	loginRouter := group.Group("/login")
	{
		loginRouter.POST("", v1.LoginService{}.Login)
	}
}
