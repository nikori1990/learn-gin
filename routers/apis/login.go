package apis

import (
	"github.com/gin-gonic/gin"
	"learn-gin/controllers/api"
)

func LoginRoutersInit(group *gin.RouterGroup) {
	loginRouter := group.Group("/login")
	{
		loginRouter.POST("", api.LoginController{}.Login)
	}
}
