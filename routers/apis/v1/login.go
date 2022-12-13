package v1

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/api/v1"
)

func LoginRoutersInit(group *gin.RouterGroup) {
	loginRouter := group.Group("/login")
	{
		loginRouter.POST("", v1.LoginController{}.Login)
	}
}
